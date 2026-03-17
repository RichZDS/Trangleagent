<template>
  <div class="home-page">
    <div class="home-bg">
      <div class="grid-overlay"></div>
      <div class="scan-line"></div>
    </div>
    <div class="logo-container">
      <div class="logo-glow"></div>
      <div class="logo-triangle">
        <div class="triangle-outer"></div>
        <div class="triangle-inner"></div>
        <div class="triangle-core"></div>
      </div>
      <div class="logo-pulse"></div>
      <p class="logo-label">TRIANGLE AGENT</p>
      <p class="logo-sublabel">SECURE TERMINAL</p>
    </div>
  </div>
</template>

<script setup lang="ts">
// 主页：三角机构 logo 展示
</script>

<style scoped>
.home-page {
  min-height: calc(100vh - 84px);
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.home-bg {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.grid-overlay {
  position: absolute;
  inset: 0;
  background-image: 
    linear-gradient(rgba(213, 0, 0, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(213, 0, 0, 0.03) 1px, transparent 1px);
  background-size: 48px 48px;
  mask-image: radial-gradient(ellipse 80% 80% at 50% 50%, black 20%, transparent 70%);
}

.scan-line {
  position: absolute;
  inset: 0;
  background: linear-gradient(
    to bottom,
    transparent 0%,
    rgba(213, 0, 0, 0.02) 50%,
    transparent 100%
  );
  background-size: 100% 8px;
  animation: scan 8s linear infinite;
  opacity: 0.6;
}

@keyframes scan {
  0% { background-position: 0 0; }
  100% { background-position: 0 100%; }
}

.logo-container {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 24px;
}

.logo-glow {
  position: absolute;
  width: 320px;
  height: 320px;
  background: radial-gradient(
    circle,
    rgba(213, 0, 0, 0.25) 0%,
    rgba(213, 0, 0, 0.08) 40%,
    transparent 70%
  );
  border-radius: 50%;
  animation: glow-pulse 4s ease-in-out infinite;
}

@keyframes glow-pulse {
  0%, 100% { transform: scale(1); opacity: 0.8; }
  50% { transform: scale(1.15); opacity: 1; }
}

.logo-triangle {
  position: relative;
  width: 200px;
  height: 173px;
}

.triangle-outer {
  position: absolute;
  left: 50%;
  top: 0;
  transform: translateX(-50%);
  width: 0;
  height: 0;
  border-left: 110px solid transparent;
  border-right: 110px solid transparent;
  border-bottom: 190px solid rgba(213, 0, 0, 0.15);
  filter: blur(0);
  animation: float 6s ease-in-out infinite;
}

.triangle-inner {
  position: absolute;
  left: 50%;
  top: 0;
  transform: translateX(-50%);
  width: 0;
  height: 0;
  border-left: 90px solid transparent;
  border-right: 90px solid transparent;
  border-bottom: 156px solid rgba(213, 0, 0, 0.4);
  filter: drop-shadow(0 0 30px rgba(213, 0, 0, 0.5));
  animation: float 6s ease-in-out infinite 0.2s;
}

.triangle-core {
  position: absolute;
  left: 50%;
  top: 0;
  transform: translateX(-50%);
  width: 0;
  height: 0;
  border-left: 70px solid transparent;
  border-right: 70px solid transparent;
  border-bottom: 121px solid #d50000;
  filter: 
    drop-shadow(0 0 20px rgba(213, 0, 0, 0.8))
    drop-shadow(0 0 60px rgba(213, 0, 0, 0.4));
  animation: float 6s ease-in-out infinite 0.4s;
}

@keyframes float {
  0%, 100% { transform: translateX(-50%) translateY(0); }
  50% { transform: translateX(-50%) translateY(-8px); }
}

.logo-pulse {
  position: absolute;
  width: 220px;
  height: 220px;
  border: 1px solid rgba(213, 0, 0, 0.2);
  border-radius: 50%;
  animation: pulse-ring 3s ease-out infinite;
  pointer-events: none;
}

.logo-pulse::before,
.logo-pulse::after {
  content: '';
  position: absolute;
  inset: -1px;
  border: 1px solid rgba(213, 0, 0, 0.15);
  border-radius: 50%;
  animation: pulse-ring 3s ease-out infinite 0.5s;
}

.logo-pulse::after {
  animation-delay: 1s;
}

@keyframes pulse-ring {
  0% {
    transform: scale(0.6);
    opacity: 1;
  }
  100% {
    transform: scale(1.4);
    opacity: 0;
  }
}

.logo-label {
  font-family: 'Roboto Mono', monospace;
  font-size: 18px;
  font-weight: 700;
  letter-spacing: 8px;
  color: var(--ta-text);
  margin: 0;
  text-shadow: 0 0 20px rgba(213, 0, 0, 0.3);
}

.logo-sublabel {
  font-family: 'Roboto Mono', monospace;
  font-size: 11px;
  color: var(--ta-text-muted);
  letter-spacing: 6px;
  margin: 0;
  opacity: 0.8;
}
</style>
