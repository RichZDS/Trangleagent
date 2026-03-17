# 多人 WebSocket 聊天室

## 技术栈

- **WebSocket**：实时双向通信
- **RabbitMQ**：多实例消息广播（fanout 交换器）

## 架构说明

1. 前端连接 `ws://host/ws/chat`，支持通过 `?token=xxx` 传递 JWT
2. 客户端发送 `join`/`leave`/`chat` 类型消息，携带 `roomId`、`userId`、`nickname`、`roleName`
3. 后端将消息发布到 RabbitMQ `chatroom_broadcast` fanout 交换器
4. 每个后端实例绑定独立队列，消费后向本实例内该房间的 WebSocket 客户端广播

## 配置

### RabbitMQ

环境变量或配置文件：

```yaml
rabbitmq:
  url: "amqp://guest:guest@localhost:5672/"
```

未配置或连接失败时，自动降级为单实例本地广播。

### 前端

- 开发环境：Vite 代理 `/ws` 到后端，使用 `ws://localhost:5173/ws/chat`
- 生产环境：通过 `VITE_WS_URL` 指定，如 `https://your-domain.com`

## 消息格式

```json
{
  "roomId": "1",
  "userId": "123",
  "nickname": "用户昵称",
  "roleName": "角色名",
  "message": "消息内容",
  "type": "join|leave|chat",
  "time": 1710000000000
}
```
