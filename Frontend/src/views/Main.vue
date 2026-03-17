<style scoped>
.main-page {
  color: #e2e8f0;
}

.hero {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  gap: 24px;
}

.hero-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.triangle-mark {
  width: 0;
  height: 0;
  border-left: 18px solid transparent;
  border-right: 18px solid transparent;
  border-bottom: 32px solid #d50000;
  filter: drop-shadow(0 0 10px rgba(213, 0, 0, 0.6));
}

.hero-text h1 {
  margin: 0;
  font-size: 20px;
  letter-spacing: 2px;
  font-family: 'Roboto Mono', monospace;
}

.hero-text p {
  margin: 4px 0 0;
  color: #94a3b8;
  font-size: 13px;
}

.hero-right {
  max-width: 360px;
  width: 100%;
}

.room-list-section {
  margin-top: 8px;
}

.section-header {
  display: flex;
  align-items: baseline;
  gap: 8px;
  margin-bottom: 16px;
}

.section-title {
  font-size: 13px;
  letter-spacing: 3px;
  color: #e5e7eb;
}

.section-sub {
  font-size: 12px;
  color: #6b7280;
}

.room-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
  gap: 16px;
}

.room-card {
  position: relative;
  background: radial-gradient(circle at 0% 0%, rgba(220, 38, 38, 0.28), rgba(15, 23, 42, 0.98));
  border-radius: 14px;
  border: 1px solid rgba(248, 113, 113, 0.4);
  padding: 16px 16px 14px;
  box-shadow: 0 18px 40px rgba(0, 0, 0, 0.65);
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.room-card-triangle {
  position: absolute;
  top: 10px;
  left: 10px;
  width: 0;
  height: 0;
  border-left: 10px solid transparent;
  border-right: 10px solid transparent;
  border-bottom: 18px solid #f97373;
  filter: drop-shadow(0 0 5px rgba(248, 113, 113, 0.7));
}

.room-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 26px;
}

.room-name {
  font-weight: 700;
  font-size: 15px;
}

.status-tag {
  border-radius: 999px;
}

.room-meta {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 6px;
  font-size: 11px;
  padding-left: 26px;
}

.meta-item {
  display: flex;
  flex-direction: column;
}

.meta-item .label {
  color: #9ca3af;
  letter-spacing: 1px;
}

.meta-item .value {
  color: #e5e7eb;
}

.room-description {
  margin-top: 4px;
  padding-left: 26px;
}

.scenario {
  display: flex;
  align-items: baseline;
  gap: 6px;
}

.scenario-label {
  font-size: 11px;
  color: #fca5a5;
  letter-spacing: 1px;
}

.scenario-name {
  font-size: 12px;
  color: #fef2f2;
}

.desc-text {
  margin: 4px 0 0;
  font-size: 12px;
  color: #e5e7eb;
  opacity: 0.85;
}

.empty-state {
  margin-top: 40px;
}

.load-more-wrap {
  margin-top: 16px;
  text-align: center;
  color: #9ca3af;
  font-size: 13px;
}

.no-more-text {
  opacity: 0.8;
}

@media (max-width: 768px) {
  .hero {
    flex-direction: column;
    align-items: flex-start;
  }

  .hero-right {
    max-width: 100%;
  }
}
</style>

<template>
  <div class="main-page">
    <div class="page-header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="page-title">房间大厅 // LOBBY</h1>
          <p class="page-subtitle">连接至活动终端</p>
        </div>
        <div class="actions">
          <a-input-search
            v-model:value="searchText"
            placeholder="搜索房间..."
            style="width: 300px"
            @search="onSearch"
            class="custom-search"
          />
          <a-button type="primary" class="create-btn" @click="showCreateModal">
            <template #icon><PlusOutlined /></template>
            新建房间
          </a-button>
        </div>
      </div>
    </div>

    <div class="room-grid-wrapper" @scroll="handleScroll">
      <div class="room-grid">
      <div v-for="room in rooms" :key="room.id" class="room-card">
        <div class="card-triangle"></div>
        <div class="card-content">
          <div class="room-status" :class="getStatusClass(room.status)">
            {{ formatStatus(room.status) }}
          </div>
          <h3 class="room-name">{{ room.roomName || room.roomCode || '未命名房间' }}</h3>
          <div class="room-info">
            <div class="info-item">
              <span class="label">剧本:</span>
              <span class="value">{{ room.scenarioName || '未设置' }}</span>
            </div>
            <div class="info-item">
              <span class="label">房主:</span>
              <span class="value">{{ room.hostNickname || ('#' + room.hostId) }}</span>
            </div>
            <div class="info-item">
              <span class="label">人数:</span>
              <span class="value">{{ room.currentPlayers || 0 }}/{{ room.maxPlayers || 0 }}</span>
            </div>
          </div>
          <div class="room-desc" v-if="room.description">
            {{ room.description }}
          </div>
          <div class="card-actions">
            <a-button type="primary" block class="join-btn" @click="openJoinModal(room)">
              接入信号
            </a-button>
          </div>
        </div>
      </div>
      <div v-if="loading" class="loading-state">
        <a-spin />
      </div>
      <div v-if="!loading && rooms.length === 0" class="empty-state">
        暂无活跃房间
      </div>
      </div>
    </div>

    <!-- 创建房间弹窗 -->
    <a-modal
      v-model:open="createModalVisible"
      title="新建终端房间"
      @ok="handleCreateRoom"
      :confirmLoading="createLoading"
      class="custom-modal"
    >
      <a-form :model="createForm" layout="vertical">
        <a-form-item label="房间名称" name="name" required>
          <a-input v-model:value="createForm.name" />
        </a-form-item>
        <a-form-item label="剧本名称" name="scenario_name" required>
          <a-input v-model:value="createForm.scenario_name" />
        </a-form-item>
        <a-form-item label="最大人数" name="max_players" required>
          <a-input-number v-model:value="createForm.max_players" :min="1" :max="10" style="width: 100%" />
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea v-model:value="createForm.description" :rows="3" />
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 加入房间：选择角色（三角机构：每个角色拥有不同素质保障） -->
    <a-modal
      v-model:open="joinModalVisible"
      title="选择角色加入房间"
      @ok="handleJoin"
      :confirmLoading="joinLoading"
      okText="接入信号"
    >
      <a-form layout="vertical">
        <a-form-item label="选择角色卡">
          <a-select
            v-model:value="selectedRoleId"
            placeholder="请选择要使用的角色"
            style="width: 100%"
            :options="joinRoles.map(r => ({ label: `${r.agentName || r.codeName || '未命名'} (${r.codeName || ''})`, value: r.id }))"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { PlusOutlined } from '@ant-design/icons-vue'
import { getApiRoomList } from '../api/controller/FangJian/getApiRoomList'
import { postApiRoomCreate } from '../api/controller/FangJian/postApiRoomCreate'
import { postApiRoomJoin } from '../api/controller/FangJian/postApiRoomJoin'
import { getApiUserView } from '../api/controller/YongHu/getApiUserView'
import { getApiRoleList } from '../api/controller/JiaoSe/getApiRoleList'

const router = useRouter()
const rooms = ref([])
const loading = ref(false)
const searchText = ref('')
const page = ref(1)
const hasMore = ref(true)

const createModalVisible = ref(false)
const createLoading = ref(false)
const createForm = reactive({
  name: '',
  scenario_name: '',
  max_players: 4,
  description: ''
})

// 加入房间：选择角色（三角机构规则：每个角色拥有不同素质保障）
const joinModalVisible = ref(false)
const joinLoading = ref(false)
const joinRoles = ref([])
const selectedRoleId = ref(null)
const joinUserId = ref(0)
const joinRoom = ref(null)

const fetchRooms = async (isLoadMore = false) => {
  if (loading.value || (!hasMore.value && isLoadMore)) return
  
  loading.value = true
  try {
    const params = {
      page: page.value,
      pageSize: 12,
    }
    if (searchText.value) {
      params.roomName = searchText.value
    }
    
    const res = await getApiRoomList(params)
    
    if (res.code && res.code !== 0) {
      message.error(res.message || '获取房间列表失败')
      return
    }
    
    const data = res.data || res || {}
    const roomList = data.list || data.rooms || []
    const total = data.total || 0
    
    if (isLoadMore) {
      rooms.value = [...rooms.value, ...roomList]
    } else {
      rooms.value = roomList
      page.value = 1
    }
    
    hasMore.value = rooms.value.length < total
    if (hasMore.value && isLoadMore) {
      page.value++
    }
  } catch (err) {
    message.error(err.message || '获取房间列表失败')
  } finally {
    loading.value = false
  }
}

const onSearch = () => {
  page.value = 1
  hasMore.value = true
  rooms.value = []
  fetchRooms()
}

const handleScroll = (e) => {
  const { scrollTop, scrollHeight, clientHeight } = e.target
  if (scrollHeight - scrollTop - clientHeight < 100 && hasMore.value && !loading.value) {
    fetchRooms(true)
  }
}

const formatStatus = (status) => {
  const map = {
    0: '待命',
    1: '执行中',
    2: '已结束',
    3: '已关闭'
  }
  return map[status] || '未知'
}

const getStatusClass = (status) => {
  const map = {
    0: 'waiting',
    1: 'playing',
    2: 'finished',
    3: 'closed'
  }
  return map[status] || ''
}

const showCreateModal = () => {
  createModalVisible.value = true
}

const handleCreateRoom = async () => {
  if (!createForm.name || !createForm.scenario_name) {
    message.warning('请填写必要信息')
    return
  }
  
  createLoading.value = true
  try {
    await postApiRoomCreate({
      roomName: createForm.name,
      scenarioName: createForm.scenario_name,
      maxPlayers: createForm.max_players,
      description: createForm.description
    })
    message.success('房间创建成功')
    createModalVisible.value = false
    onSearch() // 刷新列表
  } catch (err) {
    message.error(err.message || '创建失败')
  } finally {
    createLoading.value = false
  }
}

const openJoinModal = async (room) => {
  joinRoom.value = room
  joinRoles.value = []
  selectedRoleId.value = null
  joinUserId.value = 0

  try {
    let userId = Number(localStorage.getItem('ta_user_id'))
    if (!userId) {
      const account = localStorage.getItem('ta_account')
      if (!account) {
        message.error('请先登录')
        return
      }
      const userRes = await getApiUserView({ account })
      const userData = userRes.data || userRes || {}
      userId = userData.id
      if (userId) localStorage.setItem('ta_user_id', String(userId))
    }
    if (!userId) {
      message.error('未获取到用户信息')
      return
    }
    joinUserId.value = userId

    const roleRes = await getApiRoleList({ userId, pageSize: 50 })
    const roleData = roleRes.data || roleRes || {}
    const list = roleData.list || []
    if (list.length === 0) {
      message.warning('请先创建角色卡再加入房间')
      router.push('/roles')
      return
    }
    joinRoles.value = list
    selectedRoleId.value = list[0]?.id || null
    joinModalVisible.value = true
  } catch (err) {
    message.error(err.message || '获取角色列表失败')
  }
}

const handleJoin = async () => {
  if (!joinRoom.value || !selectedRoleId.value || !joinUserId.value) return
  joinLoading.value = true
  try {
    await postApiRoomJoin({
      roomId: joinRoom.value.id,
      userId: joinUserId.value,
      roleCardId: selectedRoleId.value
    })
    message.success(`已加入房间: ${joinRoom.value.name || joinRoom.value.roomName}`)
    joinModalVisible.value = false
    router.push({ name: 'chat-room', params: { roomId: String(joinRoom.value.id) } })
  } catch (err) {
    message.error(err?.response?.data?.message || err.message || '加入失败')
  } finally {
    joinLoading.value = false
  }
}

onMounted(() => {
  fetchRooms()
})
</script>

<style scoped>
.main-page {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 140px);
}

.page-header {
  margin-bottom: 24px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  padding-bottom: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.page-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 24px;
  color: #e2e8f0;
  margin: 0;
  letter-spacing: 2px;
}

.page-subtitle {
  color: #64748b;
  margin: 4px 0 0;
  font-size: 14px;
}

.actions {
  display: flex;
  gap: 16px;
}

.room-grid-wrapper {
  flex: 1;
  overflow-y: auto;
  padding-right: 8px;
  padding-bottom: 24px;
  max-height: calc(100vh - 240px);
}

.room-grid-wrapper::-webkit-scrollbar {
  width: 6px;
}

.room-grid-wrapper::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 3px;
}

.room-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
}

.room-card {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  position: relative;
  overflow: hidden;
  transition: all 0.3s;
  backdrop-filter: blur(4px);
}

.room-card:hover {
  transform: translateY(-4px);
  border-color: rgba(213, 0, 0, 0.5);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
}

.card-triangle {
  position: absolute;
  top: 0;
  left: 0;
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 24px 24px 0 0;
  border-color: #d50000 transparent transparent transparent;
  z-index: 2;
}

.card-content {
  padding: 20px;
}

.room-status {
  position: absolute;
  top: 12px;
  right: 12px;
  font-size: 12px;
  padding: 2px 8px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.1);
  color: #cbd5e1;
}

.room-status.playing {
  color: #22c55e;
  background: rgba(34, 197, 94, 0.1);
}

.room-status.waiting {
  color: #f59e0b;
  background: rgba(245, 158, 11, 0.1);
}

.room-status.finished {
  color: #64748b;
  background: rgba(100, 116, 139, 0.1);
}

.room-status.closed {
  color: #ef4444;
  background: rgba(239, 68, 68, 0.1);
}

.room-name {
  color: #e2e8f0;
  font-size: 18px;
  margin: 0 0 16px;
  font-family: 'Roboto Mono', monospace;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  padding-right: 60px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 14px;
}

.info-item .label {
  color: #64748b;
}

.info-item .value {
  color: #cbd5e1;
  font-weight: 500;
}

.room-desc {
  margin: 16px 0;
  color: #94a3b8;
  font-size: 13px;
  height: 40px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.join-btn {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border: none;
  font-weight: 600;
  letter-spacing: 1px;
}

.join-btn:hover {
  background: linear-gradient(120deg, #f43f5e, #991b1b);
}

.loading-state, .empty-state {
  grid-column: 1 / -1;
  text-align: center;
  padding: 40px;
  color: #64748b;
}

/* Custom UI overrides */
:deep(.ant-input-search .ant-input) {
  background: rgba(15, 23, 42, 0.6);
  border-color: rgba(255, 255, 255, 0.1);
  color: #e2e8f0;
}

:deep(.create-btn) {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.1);
}

:deep(.create-btn:hover) {
  color: #d50000;
  border-color: #d50000;
}
</style>

