<template>
  <div class="profile-page">
    <div class="page-header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="page-title">个人档案 // PROFILE</h1>
          <p class="page-subtitle">查看和编辑个人信息</p>
        </div>
        <div class="actions">
          <a-button 
            v-if="!isEditing" 
            type="primary" 
            class="action-btn edit-btn" 
            @click="startEdit"
          >
            <template #icon><EditOutlined /></template>
            编辑信息
          </a-button>
          <template v-else>
            <a-button class="action-btn cancel-btn" @click="cancelEdit">
              取消
            </a-button>
            <a-button type="primary" class="action-btn save-btn" @click="handleSave" :loading="saving">
              <template #icon><SaveOutlined /></template>
              保存
            </a-button>
          </template>
        </div>
      </div>
    </div>

    <div class="profile-content" v-if="!loading">
      <div class="profile-layout">
        <!-- 左侧：基本信息 -->
        <div class="profile-left">
          <div class="info-card">
            <div class="avatar-section">
              <div class="avatar-large">
                <img src="https://api.dicebear.com/7.x/bottts/svg?seed=agent" alt="avatar" />
              </div>
              <h2 class="user-name">{{ userInfo.nickname || userInfo.account || '未设置' }}</h2>
              <p class="user-account">ID: {{ userInfo.account || '—' }}</p>
              <!-- 经验条 -->
              <div class="exp-bar-section">
                <div class="exp-bar-header">
                  <span class="exp-level">Lv.{{ userInfo.level || 1 }}</span>
                  <span class="exp-text">{{ userInfo.exp || 0 }} EXP</span>
                </div>
                <div class="exp-bar-track">
                  <div class="exp-bar-fill" :style="{ width: expBarPercent + '%' }"></div>
                </div>
                <div class="exp-bar-hint">距离下一级还需 {{ expToNextLevel }} EXP（每级需 100 EXP）</div>
                <a-button
                  type="primary"
                  size="small"
                  class="checkin-btn"
                  :loading="checkinLoading"
                  :disabled="checkedInToday"
                  @click="handleCheckin"
                >
                  <template #icon><CalendarOutlined /></template>
                  {{ checkedInToday ? '今日已签到' : '签到 +20 EXP' }}
                </a-button>
              </div>
            </div>

            <div class="info-section">
              <h3 class="section-title">基本信息</h3>
              <div class="info-list">
                <div class="info-item">
                  <span class="info-label">账号：</span>
                  <span class="info-value">{{ userInfo.account || '—' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">昵称：</span>
                  <span class="info-value" v-if="!isEditing">{{ userInfo.nickname || '—' }}</span>
                  <a-input v-else v-model:value="form.nickname" placeholder="请输入昵称" />
                </div>
                <div class="info-item">
                  <span class="info-label">性别：</span>
                  <span class="info-value" v-if="!isEditing">{{ formatGender(userInfo.gender) }}</span>
                  <a-select v-else v-model:value="form.gender" style="width: 100%">
                    <a-select-option :value="0">未知</a-select-option>
                    <a-select-option :value="1">男</a-select-option>
                    <a-select-option :value="2">女</a-select-option>
                  </a-select>
                </div>
                <div class="info-item">
                  <span class="info-label">生日：</span>
                  <span class="info-value" v-if="!isEditing">{{ userInfo.birthDate || '—' }}</span>
                  <a-date-picker v-else v-model:value="form.birthDate" style="width: 100%" />
                </div>
                <div class="info-item">
                  <span class="info-label">用户类型：</span>
                  <span class="info-value user-type-text">
                    {{ userInfo.userType || 'normal' }}
                  </span>
                </div>
                <div class="info-item">
                  <span class="info-label">创建时间：</span>
                  <span class="info-value">{{ formatTime(userInfo.createdAt) }}</span>
                </div>
                <div class="info-item" v-if="userInfo.vipStartAt">
                  <span class="info-label">VIP开始：</span>
                  <span class="info-value">{{ formatTime(userInfo.vipStartAt) }}</span>
                </div>
                <div class="info-item" v-if="userInfo.vipEndAt">
                  <span class="info-label">VIP结束：</span>
                  <span class="info-value">{{ formatTime(userInfo.vipEndAt) }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 右侧：ARC 信息 -->
        <div class="profile-right">
          <div class="arc-card">
            <div class="arc-header">
              <h2 class="arc-title">ARC 档案</h2>
              <p class="arc-subtitle">身份识别与角色定位</p>
            </div>
            
            <div class="arc-content">
              <div class="arc-item highlight">
                <div class="arc-item-header">
                  <span class="arc-icon">🆔</span>
                  <span class="arc-label">现实身份/角色</span>
                </div>
                <div class="arc-item-content" v-if="!isEditing">
                  <span class="arc-value">{{ userInfo.realityRole || '未设置' }}</span>
                </div>
                <a-input 
                  v-else 
                  v-model:value="form.realityRole" 
                  placeholder="请输入现实身份/角色"
                  class="arc-input"
                />
              </div>

              <div class="arc-item highlight">
                <div class="arc-item-header">
                  <span class="arc-icon">⚠️</span>
                  <span class="arc-label">异常身份/角色</span>
                </div>
                <div class="arc-item-content" v-if="!isEditing">
                  <span class="arc-value">{{ userInfo.abnormalRole || '未设置' }}</span>
                </div>
                <a-input 
                  v-else 
                  v-model:value="form.abnormalRole" 
                  placeholder="请输入异常身份/角色"
                  class="arc-input"
                />
              </div>

              <div class="arc-item highlight">
                <div class="arc-item-header">
                  <span class="arc-icon">💼</span>
                  <span class="arc-label">职位</span>
                </div>
                <div class="arc-item-content" v-if="!isEditing">
                  <span class="arc-value">{{ userInfo.jobTitle || '未设置' }}</span>
                </div>
                <a-input 
                  v-else 
                  v-model:value="form.jobTitle" 
                  placeholder="请输入职位"
                  class="arc-input"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="loading-container" v-else>
      <a-spin size="large" />
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { EditOutlined, SaveOutlined, CalendarOutlined } from '@ant-design/icons-vue'
import { getApiUserView } from '../api/controller/YongHu/getApiUserView'
import { putApiUserUpdate } from '../api/controller/YongHu/putApiUserUpdate'
import { postApiUserCheckin } from '../api/controller/YongHu/postApiUserCheckin'
import dayjs from 'dayjs'

const router = useRouter()
const loading = ref(false)
const saving = ref(false)
const checkinLoading = ref(false)
const isEditing = ref(false)
const userInfo = ref({})
const form = reactive({
  id: null,
  nickname: '',
  gender: 0,
  birthDate: null,
  realityRole: '',
  abnormalRole: '',
  jobTitle: ''
})

// 经验条：每 100 EXP 升 1 级
const expBarPercent = computed(() => {
  const exp = userInfo.value?.exp || 0
  return (exp % 100)
})

const expToNextLevel = computed(() => {
  const exp = userInfo.value?.exp || 0
  return 100 - (exp % 100)
})

// 今日是否已签到
const checkedInToday = computed(() => {
  const last = userInfo.value?.lastCheckinAt
  if (!last) return false
  const lastDate = new Date(last).toDateString()
  const today = new Date().toDateString()
  return lastDate === today
})

const handleCheckin = async () => {
  const userId = userInfo.value?.id || Number(localStorage.getItem('ta_user_id'))
  if (!userId) {
    message.error('请先登录')
    return
  }
  checkinLoading.value = true
  try {
    const res = await postApiUserCheckin({ userId })
    const data = res?.data || res || {}
    userInfo.value = { ...userInfo.value, exp: data.exp, level: data.level, lastCheckinAt: new Date().toISOString() }
    message.success(data.message || '签到成功，获得 20 经验')
  } catch (err) {
    message.error(err.message || '签到失败')
  } finally {
    checkinLoading.value = false
  }
}

const fetchUserInfo = async () => {
  loading.value = true
  try {
    const storedId = localStorage.getItem('ta_user_id')
    const params = {}
    if (storedId) {
      params.id = Number(storedId)
    } else {
      const account = localStorage.getItem('ta_account')
      if (!account) {
        message.error('未找到用户信息')
        router.push('/')
        return
      }
      params.account = account
    }

    const res = await getApiUserView(params)
    
    if (res.code && res.code !== 0) {
      message.error(res.message || '获取用户信息失败')
      return
    }
    
    const data = res.data || res || {}
    if (data.id) {
      localStorage.setItem('ta_user_id', String(data.id))
    }
    userInfo.value = data
    
    Object.assign(form, {
      id: data.id,
      nickname: data.nickname || '',
      gender: data.gender || 0,
      birthDate: data.birthDate ? dayjs(data.birthDate) : null,
      realityRole: data.realityRole || '',
      abnormalRole: data.abnormalRole || '',
      jobTitle: data.jobTitle || ''
    })
  } catch (err) {
    message.error(err.message || '获取用户信息失败')
  } finally {
    loading.value = false
  }
}

const startEdit = () => {
  isEditing.value = true
}

const cancelEdit = () => {
  isEditing.value = false
  // 恢复原始数据
  Object.assign(form, {
    nickname: userInfo.value.nickname || '',
    gender: userInfo.value.gender || 0,
    birthDate: userInfo.value.birthDate ? dayjs(userInfo.value.birthDate) : null,
    realityRole: userInfo.value.realityRole || '',
    abnormalRole: userInfo.value.abnormalRole || '',
    jobTitle: userInfo.value.jobTitle || ''
  })
}

const handleSave = async () => {
  saving.value = true
  try {
    const updateData = {
      id: form.id,
      nickname: form.nickname,
      gender: form.gender,
      birthDate: form.birthDate ? (typeof form.birthDate === 'string' ? form.birthDate : form.birthDate.format('YYYY-MM-DD')) : null,
      realityRole: form.realityRole,
      abnormalRole: form.abnormalRole,
      jobTitle: form.jobTitle
    }

    await putApiUserUpdate(updateData)
    
    message.success('信息更新成功')
    isEditing.value = false
    await fetchUserInfo()
  } catch (err) {
    message.error(err.message || '更新失败')
  } finally {
    saving.value = false
  }
}

const formatGender = (gender) => {
  if (gender === 1) return '男'
  if (gender === 2) return '女'
  return '未知'
}

const formatTime = (time) => {
  if (!time) return '—'
  if (typeof time === 'string') {
    return time.replace('T', ' ').split('+')[0].split('.')[0]
  }
  return '—'
}

onMounted(() => {
  fetchUserInfo()
})
</script>

<style scoped>
.profile-page {
  min-height: 100%;
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
  gap: 12px;
}

.action-btn {
  font-family: 'Roboto Mono', monospace;
}

.edit-btn, .save-btn {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border: none;
}

.cancel-btn {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #cbd5e1;
}

.profile-content {
  margin-top: 24px;
}

.profile-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.profile-left,
.profile-right {
  display: flex;
  flex-direction: column;
}

.info-card,
.arc-card {
  background: rgba(15, 23, 42, 0.6);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  padding: 24px;
  backdrop-filter: blur(8px);
}

.avatar-section {
  text-align: center;
  padding-bottom: 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  margin-bottom: 24px;
}

.avatar-large {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  overflow: hidden;
  margin: 0 auto 16px;
  border: 3px solid rgba(213, 0, 0, 0.3);
  background: #1e293b;
}

.avatar-large img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.user-name {
  font-family: 'Roboto Mono', monospace;
  font-size: 24px;
  font-weight: 700;
  color: #e2e8f0;
  margin: 0 0 8px;
  letter-spacing: 1px;
}

.user-account {
  color: #64748b;
  font-size: 14px;
  margin: 0;
}

.exp-bar-section {
  margin-top: 16px;
  padding: 12px 16px;
  background: rgba(15, 23, 42, 0.6);
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.exp-bar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  font-size: 13px;
}

.exp-level {
  color: #f97316;
  font-weight: 600;
  font-family: 'Roboto Mono', monospace;
}

.exp-text {
  color: #94a3b8;
}

.exp-bar-track {
  height: 8px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
  overflow: hidden;
}

.exp-bar-fill {
  height: 100%;
  background: linear-gradient(90deg, #f97316, #ea580c);
  border-radius: 4px;
  transition: width 0.3s ease;
}

.exp-bar-hint {
  margin-top: 6px;
  font-size: 11px;
  color: #64748b;
  font-family: 'Roboto Mono', monospace;
}

.checkin-btn {
  margin-top: 12px;
  width: 100%;
  background: linear-gradient(120deg, #22c55e, #15803d) !important;
  border: none !important;
  font-weight: 600;
}

.checkin-btn:hover:not(:disabled) {
  background: linear-gradient(120deg, #4ade80, #22c55e) !important;
}

.checkin-btn:disabled {
  background: rgba(255, 255, 255, 0.1) !important;
  color: #64748b !important;
  cursor: not-allowed;
}

.info-section {
  margin-top: 24px;
}

.section-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 16px;
  font-weight: 700;
  color: #d50000;
  margin: 0 0 16px;
  letter-spacing: 1px;
}

.info-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 12px;
}

.info-label {
  color: #94a3b8;
  font-size: 14px;
  min-width: 80px;
  font-family: 'Roboto Mono', monospace;
}

.info-value {
  color: #e2e8f0;
  font-size: 14px;
  flex: 1;
}

.user-type-text {
  color: #ffffff;
  font-weight: 500;
}

/* ARC 卡片样式 */
.arc-header {
  text-align: center;
  padding-bottom: 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  margin-bottom: 24px;
}

.arc-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 24px;
  font-weight: 700;
  color: #d50000;
  margin: 0 0 8px;
  letter-spacing: 2px;
}

.arc-subtitle {
  color: #64748b;
  font-size: 14px;
  margin: 0;
  font-family: 'Roboto Mono', monospace;
}

.arc-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.arc-item {
  background: rgba(213, 0, 0, 0.05);
  border: 1px solid rgba(213, 0, 0, 0.2);
  border-radius: 8px;
  padding: 20px;
  transition: all 0.3s;
}

.arc-item.highlight {
  background: rgba(213, 0, 0, 0.1);
  border-color: rgba(213, 0, 0, 0.4);
  box-shadow: 0 0 20px rgba(213, 0, 0, 0.1);
}

.arc-item-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
}

.arc-icon {
  font-size: 20px;
}

.arc-label {
  font-family: 'Roboto Mono', monospace;
  font-size: 16px;
  font-weight: 700;
  color: #d50000;
  letter-spacing: 1px;
}

.arc-item-content {
  margin-top: 8px;
}

.arc-value {
  font-size: 18px;
  color: #e2e8f0;
  font-weight: 600;
  display: block;
  padding: 12px;
  background: rgba(15, 23, 42, 0.4);
  border-radius: 4px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  font-family: 'Roboto Mono', monospace;
}

.arc-input {
  margin-top: 8px;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .profile-layout {
    grid-template-columns: 1fr;
  }
}
</style>
