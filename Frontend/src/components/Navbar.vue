<template>
  <div class="nav-bar">
    <div class="brand">
      <div class="logo">
        <div class="triangle-logo"></div>
      </div>
      <div class="brand-text">
        <span class="main">TRIANGLE</span>
        <span class="sub">AGENT</span>
      </div>
    </div>
    <div class="nav-links">
      <router-link to="/main" class="nav-item" active-class="active">主页 // HOME</router-link>
      <router-link to="/rooms" class="nav-item" active-class="active">房间大厅 // LOBBY</router-link>
      <router-link v-if="isAdmin" to="/users" class="nav-item" active-class="active">人员档案 // USERS</router-link>
      <router-link to="/profile" class="nav-item" active-class="active">个人档案 // PROFILE</router-link>
      <router-link to="/roles" class="nav-item" active-class="active">特工档案 // SETTINGS</router-link>
      <router-link to="/forum" class="nav-item" active-class="active">评论 // FORUM</router-link>
    </div>
    <div class="user-profile">
      <div class="nav-exp-bar" v-if="userExp !== undefined">
        <span class="nav-exp-level">Lv.{{ userLevel || 1 }}</span>
        <div class="nav-exp-track">
          <div class="nav-exp-fill" :style="{ width: (userExp % 100) + '%' }"></div>
        </div>
      </div>
      <div class="theme-toggle" @click="toggleTheme" :title="theme === 'dark' ? '切换浅色' : '切换深色'">
        <span class="theme-icon">{{ theme === 'dark' ? '☀' : '🌙' }}</span>
      </div>
      <div class="user-info" @click="goToProfile">
        <span class="username">{{ username }}</span>
        <div class="avatar">
          <img src="https://api.dicebear.com/7.x/bottts/svg?seed=agent" alt="avatar" />
        </div>
      </div>
      <div class="logout-btn" @click="handleLogout">
        LOGOUT
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { getApiUserView } from '../api/controller/YongHu/getApiUserView'
import { useTheme } from '../composables/useTheme'

const router = useRouter()
const { theme, toggleTheme } = useTheme()
const username = ref('AGENT')
const userExp = ref(undefined)
const userLevel = ref(undefined)
const isAdmin = ref(localStorage.getItem('ta_user_type') === 'admin')

onMounted(async () => {
  const token = localStorage.getItem('ta_token')
  if (!token) {
    router.push('/')
    return
  }
  const storedAccount = localStorage.getItem('ta_account')
  if (storedAccount) {
    username.value = storedAccount
  }
  const storedId = localStorage.getItem('ta_user_id')
  if (storedId) {
    try {
      const res = await getApiUserView({ id: Number(storedId) })
      const data = res.data || res || {}
      userExp.value = data.exp ?? 0
      userLevel.value = data.level ?? 1
      if (data.userType) {
        localStorage.setItem('ta_user_type', data.userType)
        isAdmin.value = data.userType === 'admin'
      }
    } catch (_) {}
  }
})

const goToProfile = () => {
  router.push('/profile')
}

const handleLogout = () => {
  localStorage.removeItem('ta_token')
  localStorage.removeItem('ta_user_type')
  message.success('已安全登出')
  router.push('/')
}
</script>

<style scoped>
.nav-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
  padding: 0 32px;
  background: var(--ta-surface);
  border-bottom: 1px solid var(--ta-border);
  backdrop-filter: blur(10px);
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 100;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
}

.brand {
  display: flex;
  align-items: center;
  gap: 12px;
}

.triangle-logo {
  width: 0;
  height: 0;
  border-left: 10px solid transparent;
  border-right: 10px solid transparent;
  border-bottom: 18px solid #d50000;
  filter: drop-shadow(0 0 4px rgba(213, 0, 0, 0.6));
}

.brand-text {
  display: flex;
  flex-direction: column;
  line-height: 1;
}

.brand-text .main {
  font-family: 'Roboto Mono', monospace;
  font-weight: 800;
  font-size: 16px;
  color: #e2e8f0;
  letter-spacing: 1px;
}

.brand-text .sub {
  font-family: 'Roboto Mono', monospace;
  font-size: 10px;
  color: #d50000;
  letter-spacing: 2px;
}

.nav-links {
  display: flex;
  gap: 8px;
}

.nav-item {
  color: #94a3b8;
  text-decoration: none;
  font-size: 13px;
  font-weight: 600;
  padding: 8px 16px;
  border-radius: 4px;
  transition: all 0.3s;
  font-family: 'Roboto Mono', monospace;
  letter-spacing: 1px;
  position: relative;
  overflow: hidden;
}

.nav-item:hover {
  color: #e2e8f0;
  background: rgba(255, 255, 255, 0.04);
}

.nav-item.active {
  color: #d50000;
  background: rgba(213, 0, 0, 0.08);
  border: 1px solid rgba(213, 0, 0, 0.2);
}

.nav-item.active::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 2px;
  height: 100%;
  background: #d50000;
}

.theme-toggle {
  cursor: pointer;
  padding: 6px 10px;
  border-radius: 6px;
  color: #94a3b8;
  transition: all 0.2s;
}
.theme-toggle:hover {
  color: #e2e8f0;
  background: rgba(255, 255, 255, 0.08);
}
.theme-icon {
  font-size: 18px;
}

.user-profile {
  display: flex;
  align-items: center;
  gap: 16px;
}

.nav-exp-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 100px;
}

.nav-exp-level {
  font-size: 11px;
  color: #f97316;
  font-weight: 600;
  font-family: 'Roboto Mono', monospace;
  white-space: nowrap;
}

.nav-exp-track {
  flex: 1;
  height: 6px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 3px;
  overflow: hidden;
}

.nav-exp-fill {
  height: 100%;
  background: linear-gradient(90deg, #f97316, #ea580c);
  border-radius: 3px;
  transition: width 0.3s ease;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}

.user-info:hover {
  background: rgba(255, 255, 255, 0.05);
}

.username {
  color: #e2e8f0;
  font-family: 'Roboto Mono', monospace;
  font-size: 14px;
  font-weight: 600;
}

.avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: #1e293b;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.logout-btn {
  font-size: 12px;
  color: #64748b;
  cursor: pointer;
  padding: 4px 8px;
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 4px;
  transition: all 0.2s;
}

.logout-btn:hover {
  color: #ef4444;
  border-color: rgba(239, 68, 68, 0.3);
  background: rgba(239, 68, 68, 0.05);
}
</style>

