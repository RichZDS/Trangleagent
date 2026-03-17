<template>
  <div class="app-layout">
    <div class="layout-bg">
      <canvas ref="bgCanvas" class="particle-canvas"></canvas>
    </div>
    <Navbar />
    <main class="main-content">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import Navbar from './Navbar.vue'
import { useTheme } from '../composables/useTheme'

useTheme() // 初始化主题，确保 data-theme 生效

const bgCanvas = ref(null)
let ctx = null
let animationId = null
let particles = []

const resize = () => {
  if (!bgCanvas.value) return
  bgCanvas.value.width = window.innerWidth
  bgCanvas.value.height = window.innerHeight
}

const createParticles = () => {
  particles = []
  const count = 50
  for (let i = 0; i < count; i++) {
    particles.push({
      x: Math.random() * window.innerWidth,
      y: Math.random() * window.innerHeight,
      vx: (Math.random() - 0.5) * 0.2,
      vy: (Math.random() - 0.5) * 0.2,
      size: Math.random() * 2,
      alpha: Math.random() * 0.3
    })
  }
}

const animate = () => {
  if (!bgCanvas.value || !ctx) return
  ctx.clearRect(0, 0, bgCanvas.value.width, bgCanvas.value.height)
  
  particles.forEach(p => {
    p.x += p.vx
    p.y += p.vy
    
    if (p.x < 0) p.x = bgCanvas.value.width
    if (p.x > bgCanvas.value.width) p.x = 0
    if (p.y < 0) p.y = bgCanvas.value.height
    if (p.y > bgCanvas.value.height) p.y = 0
    
    ctx.fillStyle = `rgba(220, 38, 38, ${p.alpha})`
    ctx.beginPath()
    ctx.arc(p.x, p.y, p.size, 0, Math.PI * 2)
    ctx.fill()
  })
  
  // Draw connecting lines
  ctx.lineWidth = 0.5
  for (let i = 0; i < particles.length; i++) {
    for (let j = i + 1; j < particles.length; j++) {
      const dx = particles[i].x - particles[j].x
      const dy = particles[i].y - particles[j].y
      const distance = Math.sqrt(dx * dx + dy * dy)
      
      if (distance < 150) {
        ctx.strokeStyle = `rgba(220, 38, 38, ${(1 - distance / 150) * 0.1})`
        ctx.beginPath()
        ctx.moveTo(particles[i].x, particles[i].y)
        ctx.lineTo(particles[j].x, particles[j].y)
        ctx.stroke()
      }
    }
  }
  
  animationId = requestAnimationFrame(animate)
}

onMounted(() => {
  if (bgCanvas.value) {
    ctx = bgCanvas.value.getContext('2d')
    resize()
    createParticles()
    animate()
    window.addEventListener('resize', resize)
  }
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', resize)
  if (animationId) cancelAnimationFrame(animationId)
})
</script>

<style scoped>
.app-layout {
  min-height: 100vh;
  background-color: #0b1120;
  position: relative;
}

.layout-bg {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
  pointer-events: none;
  background: radial-gradient(circle at 50% 50%, var(--ta-bg-secondary) 0%, var(--ta-bg) 100%);
}

.particle-canvas {
  width: 100%;
  height: 100%;
}

.main-content {
  position: relative;
  z-index: 1;
  padding-top: 84px; /* Navbar height + padding */
  padding-bottom: 40px;
  padding-left: 20px;
  padding-right: 20px;
  max-width: 1400px;
  margin: 0 auto;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

