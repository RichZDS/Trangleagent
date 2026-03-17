package websocket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ChatMessage struct {
	RoomId  string `json:"roomId"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var clients = make(map[*websocket.Conn]bool)
var roomClients = make(map[string]map[*websocket.Conn]bool)
var mutex = sync.RWMutex{}
var once sync.Once

var mqConn *amqp.Connection
var mqChannel *amqp.Channel

func initRabbitMQ() {
	ctx := context.TODO()
	// 尝试从配置获取 RabbitMQ URL，默认为本地
	url := g.Cfg().MustGetWithEnv(ctx, "rabbitmq.url", "amqp://guest:guest@localhost:5672/").String()

	var err error
	mqConn, err = amqp.Dial(url)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ (fallback to local mode): %v", err)
		return
	}

	mqChannel, err = mqConn.Channel()
	if err != nil {
		log.Printf("Failed to open RabbitMQ channel: %v", err)
		return
	}

	err = mqChannel.ExchangeDeclare(
		"chatroom_broadcast", // name
		"fanout",             // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Printf("Failed to declare exchange: %v", err)
		return
	}

	q, err := mqChannel.QueueDeclare(
		"",    // empty name means generated temporary queue
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("Failed to declare queue: %v", err)
		return
	}

	err = mqChannel.QueueBind(
		q.Name,               // queue name
		"",                   // routing key
		"chatroom_broadcast", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to bind queue: %v", err)
		return
	}

	msgs, err := mqChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Printf("Failed to register consumer: %v", err)
		return
	}

	log.Println("RabbitMQ chatroom_broadcast connected and listening...")

	go func() {
		for d := range msgs {
			var msg ChatMessage
			if err := json.Unmarshal(d.Body, &msg); err == nil {
				broadcastLocal(msg)
			}
		}
	}()
}

func HandleChatConnections(r *ghttp.Request) {
	once.Do(func() {
		initRabbitMQ()
	})

	ws, err := upgrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
		return
	}

	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, ws)
		for roomId, members := range roomClients {
			if _, ok := members[ws]; ok {
				delete(members, ws)
				if len(members) == 0 {
					delete(roomClients, roomId)
				}
			}
		}
		mutex.Unlock()
		ws.Close()
	}()

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}

		switch msg.Type {
		case "join":
			mutex.Lock()
			members, ok := roomClients[msg.RoomId]
			if !ok {
				members = make(map[*websocket.Conn]bool)
				roomClients[msg.RoomId] = members
			}
			members[ws] = true
			mutex.Unlock()
		case "leave":
			mutex.Lock()
			if members, ok := roomClients[msg.RoomId]; ok {
				delete(members, ws)
				if len(members) == 0 {
					delete(roomClients, msg.RoomId)
				}
			}
			mutex.Unlock()
		}

		// 发布到 RabbitMQ 或降级为本地广播
		if mqChannel != nil {
			body, _ := json.Marshal(msg)
			err := mqChannel.PublishWithContext(context.Background(),
				"chatroom_broadcast", // exchange
				"",                   // routing key
				false,                // mandatory
				false,                // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			if err != nil {
				log.Printf("RabbitMQ publish error: %v", err)
				broadcastLocal(msg) // fallback
			}
		} else {
			broadcastLocal(msg) // fallback
		}
	}
}

func broadcastLocal(msg ChatMessage) {
	mutex.RLock()
	members, ok := roomClients[msg.RoomId]
	if !ok {
		mutex.RUnlock()
		return
	}
	
	// 需要关闭的客户端
	var deadClients []*websocket.Conn

	for client := range members {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("write error: %v", err)
			client.Close()
			deadClients = append(deadClients, client)
		}
	}
	mutex.RUnlock()

	// 清理断开的连接
	if len(deadClients) > 0 {
		mutex.Lock()
		for _, client := range deadClients {
			delete(members, client)
			delete(clients, client)
		}
		if len(members) == 0 {
			delete(roomClients, msg.RoomId)
		}
		mutex.Unlock()
	}
}
