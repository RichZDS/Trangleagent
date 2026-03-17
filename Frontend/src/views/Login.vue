<template>
  <div class="login-page" :class="{ scanning: loading }">
    <canvas ref="canvasRef" class="bg-canvas"></canvas>
    <div class="scan-line" v-show="loading"></div>
    <div class="login-wrap">
      <div class="badge-container">
        <div class="lanyard"></div>
        <div class="badge-inner">
          <div class="triangle-frame">
            <SafetyCertificateOutlined />
          </div>
          <div class="org-title">
            <h1>TRIANGLE</h1>
            <span>AGENCY</span>
            <h2>终端授权 // ACCESS</h2>
          </div>
          <div class="barcode">
            <div class="laser-line"></div>
          </div>
        </div>
      </div>
      
      <div class="form-panel">
        <div class="tab-switcher">
          <button 
            :class="{ active: activeKey === 'login' }"
            @click="() => activeKey = 'login'"
          >
            账号登录
          </button>
          <button 
            :class="{ active: activeKey === 'register' }"
            @click="() => activeKey = 'register'"
          >
            账号注册
          </button>
          <button 
            :class="{ active: activeKey === 'emailLogin' }"
            @click="() => activeKey = 'emailLogin'"
          >
            邮箱登录
          </button>
          <button 
            :class="{ active: activeKey === 'emailRegister' }"
            @click="() => activeKey = 'emailRegister'"
          >
            邮箱注册
          </button>
        </div>
        
        <div class="form-container">
          <div v-if="activeKey === 'login'" class="form-wrapper">
            <div class="form-header">
              <div>
                <div class="title">终端接入</div>
                <div class="subtitle">账号登录</div>
              </div>
              <div class="status-dot" />
            </div>
            <a-card bordered class="login-card">
              <a-form
                layout="vertical"
                :model="form.login"
                :rules="rules.login"
                :ref="formRefs.login"
                @finish="() => handleSubmit('login')"
              >
                <a-form-item label="特工ID" name="account" :colon="false">
                  <a-input v-model:value="form.login.account" size="large" placeholder="账号长度 5-10 位" class="id-input" />
                </a-form-item>
                <a-form-item label="密钥" name="password" :colon="false">
                  <a-input-password v-model:value="form.login.password" size="large" placeholder="请输入密码" class="id-input" />
                </a-form-item>
                <a-button type="primary" block size="large" :loading="loading" html-type="submit" class="action-button">
                  进入机构
                </a-button>
              </a-form>
            </a-card>
          </div>
          
          <div v-if="activeKey === 'register'" class="form-wrapper">
            <div class="form-header">
              <div>
                <div class="title">终端接入</div>
                <div class="subtitle">账号注册</div>
              </div>
              <div class="status-dot" />
            </div>
            <a-card bordered class="login-card">
              <a-form
                layout="vertical"
                :model="form.register"
                :rules="rules.register"
                :ref="formRefs.register"
                @finish="() => handleSubmit('register')"
              >
                <a-form-item label="特工ID" name="account" :colon="false">
                  <a-input v-model:value="form.register.account" size="large" placeholder="账号长度 5-10 位" class="id-input" />
                </a-form-item>
                <a-form-item label="密钥" name="password" :colon="false">
                  <a-input-password v-model:value="form.register.password" size="large" placeholder="至少 6 位" class="id-input" />
                </a-form-item>
                <a-form-item label="确认密钥" name="confirm" :colon="false">
                  <a-input-password v-model:value="form.register.confirm" size="large" placeholder="再次输入密码" class="id-input" />
                </a-form-item>
                <a-button type="primary" block size="large" :loading="loading" :disabled="countdown > 0" html-type="submit" class="action-button">
                  完成注册
                </a-button>
              </a-form>
            </a-card>
          </div>
          
          <div v-if="activeKey === 'emailLogin'" class="form-wrapper">
            <div class="form-header">
              <div>
                <div class="title">终端接入</div>
                <div class="subtitle">邮箱登录</div>
              </div>
              <div class="status-dot" />
            </div>
            <a-card bordered class="login-card">
              <a-form
                layout="vertical"
                :model="form.emailLogin"
                :rules="rules.email"
                :ref="formRefs.emailLogin"
                @finish="() => handleSubmit('emailLogin')"
              >
                <a-form-item label="邮箱" name="email" :colon="false">
                  <a-input v-model:value="form.emailLogin.email" size="large" placeholder="输入邮箱地址" class="id-input" />
                </a-form-item>
                <a-form-item label="验证码" name="code" :colon="false">
                  <div class="code-input-wrapper">
                    <a-input
                      v-model:value="form.emailLogin.code"
                      size="large"
                      placeholder="输入验证码"
                      class="id-input code-input"
                    />
                    <a-button
                      size="large"
                      :loading="codeLoading"
                      :disabled="countdown > 0"
                      @click="sendCode"
                      class="code-send-btn"
                    >
                      {{ codeBtnText }}
                    </a-button>
                  </div>
                </a-form-item>
                <a-button type="primary" block size="large" :loading="loading" html-type="submit" class="action-button">
                  验证并登录
                </a-button>
              </a-form>
            </a-card>
          </div>
          
          <div v-if="activeKey === 'emailRegister'" class="form-wrapper">
            <div class="form-header">
              <div>
                <div class="title">终端接入</div>
                <div class="subtitle">邮箱注册</div>
              </div>
              <div class="status-dot" />
            </div>
            <a-card bordered class="login-card">
              <a-form
                layout="vertical"
                :model="form.emailRegister"
                :rules="rules.email"
                :ref="formRefs.emailRegister"
                @finish="() => handleSubmit('emailRegister')"
              >
                <a-form-item label="邮箱" name="email" :colon="false">
                  <a-input v-model:value="form.emailRegister.email" size="large" placeholder="输入邮箱地址" class="id-input" />
                </a-form-item>
                <a-form-item label="验证码" name="code" :colon="false">
                  <div class="code-input-wrapper">
                    <a-input
                      v-model:value="form.emailRegister.code"
                      size="large"
                      placeholder="输入验证码"
                      class="id-input code-input"
                    />
                    <a-button
                      size="large"
                      :loading="codeLoading"
                      :disabled="countdown > 0"
                      @click="sendCode"
                      class="code-send-btn"
                    >
                      {{ codeBtnText }}
                    </a-button>
                  </div>
                </a-form-item>
                <a-button type="primary" block size="large" :loading="loading" html-type="submit" class="action-button">
                  完成注册
                </a-button>
              </a-form>
            </a-card>
          </div>
          
          <div class="foot-note">端口 8000 / 代理已启用 / TOKEN 将本地缓存</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onBeforeUnmount, onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { SafetyCertificateOutlined } from '@ant-design/icons-vue'
import { postApiLogin } from '../api/controller/DengLu/postApiLogin'
import { postApiRegister } from '../api/controller/ZhuCe/postApiRegister'
import { postApiLoginEmail } from '../api/controller/DengLu/postApiLoginEmail'
import { postApiRegisterEmail } from '../api/controller/ZhuCe/postApiRegisterEmail'
import { postApiEmailSendCode } from '../api/controller/YouXiang/postApiEmailSendCode'

const router = useRouter()
const activeKey = ref('login')
const loading = ref(false)
const codeLoading = ref(false)
const countdown = ref(0)
let timer = null
const canvasRef = ref(null)
let rafId = 0
let particles = []
let ctx = null
let width = 0
let height = 0

const form = reactive({
  login: {
    account: '',
    password: '',
  },
  register: {
    account: '',
    password: '',
    confirm: '',
  },
  emailLogin: {
    email: '',
    code: '',
  },
  emailRegister: {
    email: '',
    code: '',
  },
})

const formRefs = {
  login: ref(null),
  register: ref(null),
  emailLogin: ref(null),
  emailRegister: ref(null),
}

const rules = {
  login: {
    account: [{ required: true, message: '请输入账号' }],
    password: [
      { required: true, message: '请输入密码' },
      { min: 6, message: '密码长度不能少于6个字符' },
      { max: 32, message: '密码长度不能超过32个字符' }
    ],
  },
  register: {
    account: [{ required: true, message: '请输入账号' }],
    password: [
      { required: true, message: '请输入密码' },
      { min: 6, message: '密码长度不能少于6个字符' },
      { max: 32, message: '密码长度不能超过32个字符' }
    ],
    confirm: [{ required: true, message: '请确认密码' }],
  },
  email: {
    email: [
      { required: true, message: '请输入邮箱' },
      { type: 'email', message: '邮箱格式不正确' },
    ],
    code: [{ required: true, message: '请输入验证码' }],
  },
}

const codeBtnText = computed(() =>
  countdown.value > 0 ? `${countdown.value}s` : '获取验证码',
)

const pickEmail = () =>
  activeKey.value === 'emailRegister'
    ? form.emailRegister.email
    : form.emailLogin.email

const saveToken = (token, userType) => {
  if (token) {
    localStorage.setItem('ta_token', token)
  }
  if (userType) {
    localStorage.setItem('ta_user_type', userType)
  }
}

const handleSubmit = async (type) => {
  const refItem = formRefs[type]?.value
  if (!refItem) return
  await refItem.validate()
  if (type === 'register' && form.register.password !== form.register.confirm) {
    message.error('两次输入的密码不一致')
    return
  }
  
  // 显示扫描动画
  loading.value = true
  
  // 等待扫描动画完成后再执行实际操作
  setTimeout(async () => {
    try {
      if (type === 'login') {
        const res = await postApiLogin(form.login)
        if (res.code && res.code !== 0) {
          message.error(res.message || '登录失败')
          loading.value = false
          return
        }
        const data = res.data || res || {}
        saveToken(data.token, data.userType)
        if (data.account) {
          localStorage.setItem('ta_account', data.account)
        }
        message.success('登录成功')
        router.push('/main')
      } else if (type === 'register') {
        const res = await postApiRegister({
          account: form.register.account,
          password: form.register.password,
        })
        if (res.code && res.code !== 0) {
          message.error(res.message || '注册失败')
          loading.value = false
          return
        }
        message.success('注册成功，请登录')
        activeKey.value = 'login'
        loading.value = false
      } else if (type === 'emailLogin') {
        const res = await postApiLoginEmail(form.emailLogin)
        if (res.code && res.code !== 0) {
          message.error(res.message || '邮箱登录失败')
          loading.value = false
          return
        }
        const data = res.data || res || {}
        saveToken(data.token, data.userType)
        if (data.account || data.email) {
          localStorage.setItem('ta_account', data.account || data.email)
        }
        message.success('邮箱登录成功')
        router.push('/main')
      } else if (type === 'emailRegister') {
        const res = await postApiRegisterEmail(form.emailRegister)
        if (res.code && res.code !== 0) {
          message.error(res.message || '邮箱注册失败')
          loading.value = false
          return
        }
        message.success('邮箱注册成功，请登录')
        activeKey.value = 'emailLogin'
        loading.value = false
      }
    } catch (err) {
      message.error(err.message || '请求失败')
      loading.value = false
    }
  }, 2000)
}

const sendCode = async () => {
  const email = pickEmail()
  if (!email) {
    message.warning('请先填写邮箱')
    return
  }
  codeLoading.value = true
  try {
    await postApiEmailSendCode({ email })
    message.success('验证码已发送')
    startCountdown()
  } catch (err) {
    message.error(err.message || '发送失败')
  } finally {
    codeLoading.value = false
  }
}

const startCountdown = () => {
  countdown.value = 60
  clearInterval(timer)
  timer = setInterval(() => {
    countdown.value -= 1
    if (countdown.value <= 0) {
      clearInterval(timer)
    }
  }, 1000)
}

onBeforeUnmount(() => {
  clearInterval(timer)
  cancelAnimationFrame(rafId)
  window.removeEventListener('resize', resizeCanvas)
})

const createParticles = () => {
  particles = Array.from({ length: 70 }).map(() => ({
    x: Math.random() * width,
    y: Math.random() * height,
    vx: (Math.random() - 0.5) * 0.4,
    vy: (Math.random() - 0.5) * 0.4,
    size: Math.random() * 2 + 0.5,
  }))
}

const resizeCanvas = () => {
  if (!canvasRef.value) return
  width = canvasRef.value.width = window.innerWidth
  height = canvasRef.value.height = window.innerHeight
}

const drawParticles = () => {
  if (!ctx) return
  ctx.clearRect(0, 0, width, height)
  particles.forEach((p) => {
    p.x += p.vx
    p.y += p.vy
    if (p.x < 0 || p.x > width) p.vx *= -1
    if (p.y < 0 || p.y > height) p.vy *= -1
    ctx.fillStyle = 'rgba(255,255,255,0.35)'
    ctx.beginPath()
    ctx.arc(p.x, p.y, p.size, 0, Math.PI * 2)
    ctx.fill()
  })
  for (let i = 0; i < particles.length; i++) {
    for (let j = i + 1; j < particles.length; j++) {
      const a = particles[i]
      const b = particles[j]
      const dist = Math.hypot(a.x - b.x, a.y - b.y)
      if (dist < 120) {
        ctx.strokeStyle = `rgba(255,255,255,${0.12 - dist / 120 * 0.12})`
        ctx.beginPath()
        ctx.moveTo(a.x, a.y)
        ctx.lineTo(b.x, b.y)
        ctx.stroke()
      }
    }
  }
  rafId = requestAnimationFrame(drawParticles)
}

onMounted(() => {
  if (canvasRef.value) {
    ctx = canvasRef.value.getContext('2d')
    resizeCanvas()
    createParticles()
    drawParticles()
    window.addEventListener('resize', () => {
      resizeCanvas()
      createParticles()
    })
  }
})
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #0b1120;
  position: relative;
  padding: 32px 16px;
  color: #e2e8f0;
  overflow: hidden;
}
.login-page.scanning::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(to bottom, transparent 0%, rgba(213, 0, 0, 0.1) 50%, transparent 100%);
  z-index: 100;
  animation: scanPage 2s linear;
  pointer-events: none;
}
.scan-line {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  background: #d50000;
  box-shadow: 0 0 10px #d50000, 0 0 20px #d50000;
  animation: scanPage 2s linear;
  z-index: 101;
  pointer-events: none;
}
@keyframes scanPage {
  0% {
    transform: translateY(-100%);
  }
  100% {
    transform: translateY(100vh);
  }
}
.bg-canvas {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}
.login-wrap {
  position: relative;
  z-index: 2;
  width: min(550px, 100%);
  display: flex;
  flex-direction: column;
  gap: 24px;
  align-items: stretch;
  margin: 0 auto;
}
.badge-container {
  position: relative;
  background: linear-gradient(90deg, #1e293b 0%, #0f172a 100%);
  border-radius: 14px;
  box-shadow: 0 25px 60px rgba(0, 0, 0, 0.5);
  overflow: hidden;
  min-height: 200px;
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 20px;
}
.lanyard {
  position: absolute;
  top: 0;
  width: 100%;
  height: 18px;
  background: linear-gradient(180deg, #0f172a, #1f2937, #0f172a);
  left: 0;
  opacity: 0.8;
}
.badge-inner {
  position: relative;
  z-index: 2;
  width: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 30px;
}
.triangle-frame {
  width: 90px;
  height: 90px;
  background: #d50000;
  clip-path: polygon(50% 0%, 0% 100%, 100% 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 36px;
  box-shadow: 0 8px 15px rgba(213, 0, 0, 0.3);
  flex-shrink: 0;
}
.org-title {
  text-align: left;
  line-height: 1;
}
.org-title h1 {
  font-family: 'Roboto Mono', monospace;
  letter-spacing: 2px;
  margin: 0;
  font-size: 26px;
  color: #f1f5f9;
}
.org-title span {
  display: block;
  font-family: 'Roboto Mono', monospace;
  font-size: 22px;
  color: #f87171;
  letter-spacing: 2px;
  margin-top: 2px;
}
.org-title h2 {
  margin: 10px 0 0;
  font-size: 10px;
  color: #94a3b8;
  letter-spacing: 4px;
}
.barcode {
  width: 200px;
  height: 100%;
  background: repeating-linear-gradient(0deg, #334155 0, #334155 4px, transparent 4px, transparent 8px);
  position: relative;
  opacity: 0.6;
  flex-shrink: 0;
}
.laser-line {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  background: #d50000;
  box-shadow: 0 0 15px #d50000;
  animation: scan 1.5s infinite ease-in-out;
}
@keyframes scan {
  0% {
    top: 0;
    opacity: 0;
  }
  15% {
    opacity: 1;
  }
  85% {
    opacity: 1;
  }
  100% {
    top: 100%;
    opacity: 0;
  }
}
.form-panel {
  flex: 1;
  background: rgba(30, 41, 59, 0.6);
  border-radius: 12px;
  padding: 16px;
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255, 255, 255, 0.15);
  display: flex;
  flex-direction: row;
  gap: 20px;
}
.form-wrapper {
  width: 100%;
}
.form-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
  color: #e2e8f0;
}
.title {
  font-size: 18px;
  font-weight: 800;
  color: #e2e8f0;
}
.subtitle {
  color: #94a3b8;
  font-size: 13px;
  margin-top: 4px;
}
.status-dot {
  width: 12px;
  height: 12px;
  background: #22c55e;
  border-radius: 50%;
  box-shadow: 0 0 12px #22c55e;
}
.foot-note {
  margin-top: 16px;
  color: #94a3b8;
  font-size: 12px;
  text-align: right;
}
.tab-switcher {
  width: 120px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.tab-switcher button {
  background: rgba(15, 23, 42, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #cbd5e1;
  padding: 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  text-align: left;
  font-size: 14px;
}
.tab-switcher button.active {
  background: rgba(213, 0, 0, 0.3);
  color: #ffffff;
  font-weight: 700;
  border-color: rgba(213, 0, 0, 0.5);
}
.tab-switcher button:hover:not(.active) {
  background: rgba(15, 23, 42, 0.6);
}
.form-container {
  flex: 1;
}
.code-input-wrapper {
  display: flex;
  gap: 8px;
  width: 100%;
}

.code-input {
  flex: 1;
}

.code-send-btn {
  flex-shrink: 0;
  min-width: 120px;
  background: rgba(30, 41, 59, 0.6) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  color: #d50000 !important;
}

.code-send-btn:hover:not(:disabled) {
  background: rgba(30, 41, 59, 0.8) !important;
  border-color: rgba(213, 0, 0, 0.5) !important;
  color: #ff1744 !important;
}

.code-send-btn:disabled {
  background: rgba(30, 41, 59, 0.4) !important;
  border-color: rgba(255, 255, 255, 0.1) !important;
  color: rgba(213, 0, 0, 0.5) !important;
  cursor: not-allowed;
}

.id-input {
  width: 100% !important;
  min-height: 48px;
}
:deep(.ant-form-item-label label) {
  color: #d50000 !important;
  font-weight: bold;
}
:deep(.ant-form-item-required) {
  padding-left: 0 !important;
}
:deep(.ant-form-item-required::before) {
  content: '' !important;
  margin-right: 0 !important;
}
.login-card {
  border-radius: 12px;
  background: rgba(15, 23, 42, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.08);
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.35);
}
.ant-card-head {
  border-bottom: none;
  background: transparent;
}
:deep(.ant-card-body) {
  background: transparent;
  color: #e2e8f0;
}
:deep(.ant-input),
:deep(.ant-input-password input),
:deep(.ant-input-affix-wrapper),
:deep(.ant-input-search) {
  background: rgba(255, 255, 255, 0.06) !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
  color: #e2e8f0 !important;
  width: 100% !important;
}
:deep(.ant-input-search .ant-input) {
  background: transparent !important;
}
:deep(.ant-input::placeholder),
:deep(.ant-input-password input::placeholder) {
  color: #94a3b8;
}
:deep(.ant-btn-primary) {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border-color: transparent;
}
:deep(.ant-btn-primary:hover) {
  background: linear-gradient(120deg, #f43f5e, #991b1b);
}
.action-button {
  max-width: 200px;
  margin: 0 auto;
}
@media (max-width: 768px) {
  .badge-container {
    flex-direction: column;
    text-align: center;
    min-height: auto;
    padding: 20px 15px;
  }
  .badge-inner {
    flex-direction: column;
    gap: 15px;
    width: 100%;
  }
  .org-title {
    text-align: center;
  }
  .form-panel {
    flex-direction: column;
  }
  .tab-switcher {
    width: 100%;
    flex-direction: row;
    flex-wrap: wrap;
  }
  .barcode {
    width: 100%;
    height: 30px;
  }
  .laser-line {
    width: 100%;
    height: 3px;
  }
}
</style>