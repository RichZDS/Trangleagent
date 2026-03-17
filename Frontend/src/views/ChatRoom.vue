<template>
  <div class="chat-room-container">
    <div class="chat-header">
      <h2>房间连接中继站: {{ roomId }}</h2>
      <a-button type="primary" danger @click="handleLeave">断开连接</a-button>
    </div>

    <div class="chat-messages" ref="messagesContainer">
      <div v-if="messages.length === 0" class="empty-state">
        暂无通讯记录，等待信号接入...
      </div>
      <div 
        v-for="(msg, index) in messages" 
        :key="index"
        :class="['message-item', msg.userId === currentUserId ? 'message-self' : 'message-other']"
      >
        <div class="message-sender">{{ msg.userId }}</div>
        <div class="message-content">{{ msg.message }}</div>
      </div>
    </div>

    <div class="chat-input-area">
      <a-input
        v-model:value="inputText"
        placeholder="输入要发送的信息..."
        @pressEnter="sendMessage"
        class="custom-input"
      />
      <a-button type="primary" @click="sendMessage" class="send-btn" :disabled="!wsConnected">
        发送信号
      </a-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { message } from 'ant-design-vue'

const route = useRoute()
const router = useRouter()

const roomId = ref(route.params.roomId || route.query.roomId || '未知房间')
// 为了简单测试，从 localStorage 拿 ID，如果为空可以生成随机的
const userStr = localStorage.getItem('ta_user')
const currentUserId = ref(userStr ? JSON.parse(userStr).id : 'User_' + Math.floor(Math.random() * 1000))

const messages = ref([])
const inputText = ref('')
const messagesContainer = ref(null)

let ws = null
const wsConnected = ref(false)

const connectWebSocket = () => {
  const wsUrl = `ws://localhost:8888/ws/chat`
  ws = new WebSocket(wsUrl)

  ws.onopen = () => {
    wsConnected.value = true
    message.success('已连接至通讯网络')
    
    // 发送加入房间消息
    ws.send(JSON.stringify({
      roomId: roomId.value,
      userId: String(currentUserId.value),
      message: '进入了频道',
      type: 'join'
    }))
  }

  ws.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data)
      if (data.roomId === roomId.value) {
        messages.value.push(data)
        scrollToBottom()
      }
    } catch (e) {
      console.error('解析消息失败', e)
    }
  }

  ws.onclose = () => {
    wsConnected.value = false
    message.warning('通讯网络已断开')
  }

  ws.onerror = (error) => {
    console.error('WebSocket Error', error)
    message.error('通讯网络连接错误')
  }
}

const sendMessage = () => {
  if (!inputText.value.trim() || !wsConnected.value) return

  const msgPayload = {
    roomId: String(roomId.value),
    userId: String(currentUserId.value),
    message: inputText.value,
    type: 'chat'
  }

  ws.send(JSON.stringify(msgPayload))
  inputText.value = ''
}

const handleLeave = () => {
  if (ws && wsConnected.value) {
    ws.send(JSON.stringify({
      roomId: String(roomId.value),
      userId: String(currentUserId.value),
      message: '离开了频道',
      type: 'leave'
    }))
    ws.close()
  }
  router.push('/rooms')
}

const scrollToBottom = () => {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
    }
  })
}

onMounted(() => {
  if (!roomId.value || roomId.value === '未知房间') {
    message.error('无效的房间号')
    router.push('/rooms')
    return
  }
  connectWebSocket()
})

onUnmounted(() => {
  if (ws) {
    ws.close()
  }
})
</script>

<style scoped>
.chat-room-container {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 140px);
  background: rgba(15, 23, 42, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  overflow: hidden;
}

.chat-header {
  padding: 16px 24px;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.chat-header h2 {
  color: #e2e8f0;
  margin: 0;
  font-family: 'Roboto Mono', monospace;
  font-size: 18px;
}

.chat-messages {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.empty-state {
  color: #64748b;
  text-align: center;
  margin-top: 40px;
}

.message-item {
  display: flex;
  flex-direction: column;
  max-width: 70%;
}

.message-self {
  align-self: flex-end;
  align-items: flex-end;
}

.message-other {
  align-self: flex-start;
  align-items: flex-start;
}

.message-sender {
  font-size: 12px;
  color: #94a3b8;
  margin-bottom: 4px;
}

.message-content {
  padding: 10px 16px;
  border-radius: 8px;
  font-size: 14px;
  line-height: 1.5;
  word-break: break-word;
}

.message-self .message-content {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  color: #fff;
  border-bottom-right-radius: 0;
}

.message-other .message-content {
  background: rgba(30, 41, 59, 0.8);
  color: #e2e8f0;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-bottom-left-radius: 0;
}

.chat-input-area {
  padding: 16px 24px;
  background: rgba(0, 0, 0, 0.3);
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  gap: 16px;
}

.custom-input {
  background: rgba(15, 23, 42, 0.6);
  border-color: rgba(255, 255, 255, 0.1);
  color: #e2e8f0;
}

.send-btn {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border: none;
  font-weight: 600;
  letter-spacing: 1px;
}

.send-btn:hover {
  background: linear-gradient(120deg, #f43f5e, #991b1b);
}
</style>
