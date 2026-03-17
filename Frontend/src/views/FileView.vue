<template>
  <div class="file-view-container" :class="trackColorClass">
    <!-- 阶段1: Loading动画 -->
    <div v-if="stage === 'loading'" class="loading-stage">
      <div class="triangle-loader" :style="{ borderBottomColor: trackColor }"></div>
      <div class="loading-text" :style="{ color: trackColor }">LOADING...</div>
    </div>

    <!-- 阶段2: 扫描动画 -->
    <div v-if="stage === 'scanning'" class="scanning-stage">
      <div class="scan-container">
        <div class="scan-line" :style="{ background: `linear-gradient(90deg, transparent, ${trackColor}, transparent)`, boxShadow: `0 0 10px ${trackColor}` }"></div>
      </div>
    </div>

    <!-- 阶段3: 验证成功 -->
    <div v-if="stage === 'verified'" class="verified-stage">
      <div class="verification-text" :style="{ color: '#4ade80' }">验证成功，符合要求</div>
    </div>

    <!-- 阶段4: 文件袋动画 -->
    <div v-if="stage === 'folder'" class="folder-stage">
      <div class="folder-container">
        <!-- 文件袋背景 -->
        <div class="folder-background">
          <div class="folder-seal">
            <div class="triangle-logo-large" :style="{ borderBottomColor: trackColor }"></div>
            <div class="agency-logo">
              <div class="logo-text" :style="{ color: trackColor }">TRIANGLE</div>
              <div class="logo-subtext" :style="{ color: trackColor }">AGENCY</div>
            </div>
          </div>
        </div>
        
        <!-- 文件袋盖子 -->
        <div class="folder-lid" :class="{ 'open': folderOpen }">
          <div class="lid-inner"></div>
        </div>

        <!-- 文件 -->
        <div 
          class="file-document" 
          :class="{ 'extracted': fileExtracted }" 
          :style="extractedStyle"
        >
          <div 
            class="file-content" 
            @wheel.prevent="handleWheel" 
            @mousedown="startDrag"
            :style="contentStyle"
          >
            <img 
              :src="imageUrl" 
              :alt="filename"
              class="file-image"
              :style="imageStyle"
              @load="onImageLoad"
              @error="onImageError"
              draggable="false"
            />
            <div class="zoom-controls" @mousedown.stop>
              <button class="zoom-btn" @click.stop="zoomIn">+</button>
              <button class="zoom-btn" @click.stop="zoomOut">−</button>
              <button class="zoom-btn" @click.stop="resetZoom">↺</button>
            </div>
            <button class="close-btn" @click.stop="closeFile" @mousedown.stop>
              <span>×</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

const stage = ref<'loading' | 'scanning' | 'verified' | 'folder'>('loading');
const folderOpen = ref(false);
const fileExtracted = ref(false);
const imageLoaded = ref(false);
const imageScale = ref(1);
const isDragging = ref(false);
const dragStart = ref({ x: 0, y: 0 });
const imagePosition = ref({ x: 0, y: 0 });

const trackType = ref(route.query.trackType as string || 'blue');
const filename = ref(route.query.filename as string || '');

const trackColor = computed(() => {
  switch(trackType.value) {
    case 'blue': return '#3498db';
    case 'yellow': return '#f1c40f';
    case 'red': return '#e74c3c';
    default: return '#3498db';
  }
});

const trackColorClass = computed(() => {
  return `track-${trackType.value}`;
});

const imageUrl = computed(() => {
  const trackName = trackType.value.charAt(0).toUpperCase() + trackType.value.slice(1);
  return `/File/${trackName}/${filename.value}.png`;
});

// 计算样式
const extractedStyle = computed(() => {
  if (!fileExtracted.value) {
    return {
      transform: 'translate(-50%, -50%) scale(0.2) translateZ(-300px) rotateY(180deg)',
      opacity: '0'
    };
  }
  return {
    transform: 'translate(-50%, -50%) translateZ(0) rotateY(0deg)',
    opacity: '1'
  };
});

const contentStyle = computed(() => {
  return {
    transform: `translate(${imagePosition.value.x}px, ${imagePosition.value.y}px)`,
    cursor: isDragging.value ? 'grabbing' : 'grab'
  };
});

const imageStyle = computed(() => {
  return {
    transform: `scale(${imageScale.value})`,
    transformOrigin: 'center center'
  };
});

// 动画序列
onMounted(() => {
  // 阶段1: Loading (1.5秒)
  setTimeout(() => {
    stage.value = 'scanning';
  }, 1500);

  // 阶段2: 扫描 (2秒)
  setTimeout(() => {
    stage.value = 'verified';
  }, 3500);

  // 阶段3: 验证成功 (1.5秒)
  setTimeout(() => {
    stage.value = 'folder';
    // 打开文件袋
    setTimeout(() => {
      folderOpen.value = true;
      // 拿出文件
      setTimeout(() => {
        fileExtracted.value = true;
      }, 1000);
    }, 300);
  }, 5000);
});

const onImageLoad = () => {
  imageLoaded.value = true;
};

const onImageError = () => {
  console.error('Failed to load image:', imageUrl.value);
};

// 缩放控制
const zoomIn = () => {
  imageScale.value = Math.min(imageScale.value + 0.2, 3);
};

const zoomOut = () => {
  imageScale.value = Math.max(imageScale.value - 0.2, 0.5);
};

const resetZoom = () => {
  imageScale.value = 1;
  imagePosition.value = { x: 0, y: 0 };
};

// 鼠标滚轮缩放
const handleWheel = (e: WheelEvent) => {
  const delta = e.deltaY > 0 ? -0.1 : 0.1;
  imageScale.value = Math.max(0.5, Math.min(3, imageScale.value + delta));
};

// 拖拽功能
const startDrag = (e: MouseEvent) => {
  if (e.button === 0 && fileExtracted.value) { // 左键且文件已提取
    e.preventDefault();
    isDragging.value = true;
    dragStart.value = { x: e.clientX - imagePosition.value.x, y: e.clientY - imagePosition.value.y };
    document.addEventListener('mousemove', handleDrag);
    document.addEventListener('mouseup', endDrag);
  }
};

const handleDrag = (e: MouseEvent) => {
  if (isDragging.value) {
    imagePosition.value = {
      x: e.clientX - dragStart.value.x,
      y: e.clientY - dragStart.value.y
    };
  }
};

const endDrag = () => {
  if (isDragging.value) {
    isDragging.value = false;
    document.removeEventListener('mousemove', handleDrag);
    document.removeEventListener('mouseup', endDrag);
  }
};

const closeFile = () => {
  window.close();
  // 如果无法关闭（某些浏览器限制），则返回上一页
  if (window.history.length > 1) {
    router.back();
  } else {
    router.push('/roles');
  }
};
</script>

<style scoped>
.file-view-container {
  width: 100vw;
  height: 100vh;
  background: #0f172a;
  position: relative;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Loading阶段 */
.loading-stage {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 40px;
}

.triangle-loader {
  width: 0;
  height: 0;
  border-left: 30px solid transparent;
  border-right: 30px solid transparent;
  border-bottom: 50px solid #d50000;
  animation: pulse 1.5s infinite;
}

.loading-text {
  font-family: 'Roboto Mono', monospace;
  font-size: 18px;
  font-weight: 700;
  letter-spacing: 4px;
  color: #d50000;
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.3);
    opacity: 0.7;
  }
}

/* 扫描阶段 */
.scanning-stage {
  width: 100%;
  height: 100%;
  position: relative;
}

.scan-container {
  width: 100%;
  height: 100%;
  position: relative;
  overflow: hidden;
}

.scan-line {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  animation: scanDown 2s linear forwards;
  transform: translateY(-100%);
}

@keyframes scanDown {
  0% {
    transform: translateY(-100%);
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  90% {
    opacity: 1;
  }
  100% {
    transform: translateY(100vh);
    opacity: 0;
  }
}

/* 验证成功阶段 */
.verified-stage {
  display: flex;
  align-items: center;
  justify-content: center;
}

.verification-text {
  font-family: 'Roboto Mono', monospace;
  font-size: 20px;
  font-weight: 700;
  letter-spacing: 2px;
  color: #4ade80;
  animation: fadeInUp 0.5s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 文件袋阶段 */
.folder-stage {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
}

.folder-container {
  position: relative;
  width: 600px;
  height: 800px;
  perspective: 1000px;
}

.folder-background {
  position: absolute;
  width: 100%;
  height: 100%;
  background: 
    radial-gradient(circle at 20% 30%, rgba(0, 0, 0, 0.3) 0%, transparent 50%),
    linear-gradient(135deg, #1a252f 0%, #2c3e50 100%);
  border-radius: 12px;
  box-shadow: 
    0 20px 60px rgba(0, 0, 0, 0.8),
    inset 0 0 100px rgba(0, 0, 0, 0.3);
  border: 2px solid rgba(255, 255, 255, 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  position: relative;
}

.folder-background::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    repeating-linear-gradient(
      0deg,
      transparent,
      transparent 2px,
      rgba(0, 0, 0, 0.1) 2px,
      rgba(0, 0, 0, 0.1) 4px
    ),
    repeating-linear-gradient(
      90deg,
      transparent,
      transparent 2px,
      rgba(0, 0, 0, 0.1) 2px,
      rgba(0, 0, 0, 0.1) 4px
    );
  opacity: 0.3;
  pointer-events: none;
}

.folder-seal {
  text-align: center;
  z-index: 2;
  position: relative;
  padding: 40px;
  background: radial-gradient(circle at center, rgba(0, 0, 0, 0.4) 0%, transparent 70%);
  border-radius: 50%;
}

.triangle-logo-large {
  width: 0;
  height: 0;
  border-left: 60px solid transparent;
  border-right: 60px solid transparent;
  border-bottom: 100px solid;
  margin: 0 auto 20px;
  filter: drop-shadow(0 0 20px currentColor);
  animation: glow 2s ease-in-out infinite;
  position: relative;
}

.triangle-logo-large::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 0;
  height: 0;
  border-left: 20px solid transparent;
  border-right: 20px solid transparent;
  border-bottom: 35px solid rgba(255, 255, 255, 0.2);
}

@keyframes glow {
  0%, 100% {
    filter: drop-shadow(0 0 20px currentColor);
    transform: scale(1);
  }
  50% {
    filter: drop-shadow(0 0 40px currentColor);
    transform: scale(1.05);
  }
}

.agency-logo {
  margin-top: 20px;
}

.logo-text {
  font-family: 'Roboto Mono', monospace;
  font-size: 32px;
  font-weight: 800;
  letter-spacing: 4px;
  margin-bottom: 8px;
}

.logo-subtext {
  font-family: 'Roboto Mono', monospace;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 6px;
  opacity: 0.8;
}

.folder-lid {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: 
    radial-gradient(circle at 20% 30%, rgba(0, 0, 0, 0.4) 0%, transparent 50%),
    linear-gradient(135deg, #1a252f 0%, #2c3e50 100%);
  border-radius: 12px;
  transform-origin: top center;
  transform: rotateX(0deg);
  transition: transform 0.8s cubic-bezier(0.25, 1, 0.5, 1);
  z-index: 3;
  box-shadow: 
    0 -10px 30px rgba(0, 0, 0, 0.5),
    inset 0 0 100px rgba(0, 0, 0, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.1);
  border-bottom: none;
}

.folder-lid.open {
  transform: rotateX(-120deg);
}

.lid-inner {
  width: 100%;
  height: 100%;
  background: repeating-linear-gradient(
    45deg,
    rgba(0, 0, 0, 0.1) 0px,
    rgba(0, 0, 0, 0.1) 10px,
    transparent 10px,
    transparent 20px
  );
  border-radius: 12px;
}

.file-document {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%) scale(0.2) translateZ(-300px) rotateY(180deg);
  opacity: 0;
  transition: all 1.2s cubic-bezier(0.25, 1, 0.5, 1);
  z-index: 1;
  pointer-events: none;
}

.file-document.extracted {
  opacity: 1;
  pointer-events: all;
  z-index: 10;
  transition: opacity 1.2s cubic-bezier(0.25, 1, 0.5, 1);
}

.file-content {
  position: relative;
  background: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 
    0 20px 60px rgba(0, 0, 0, 0.6),
    0 0 0 1px rgba(0, 0, 0, 0.1);
  width: 560px;
  height: 760px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  cursor: grab;
  user-select: none;
}

.file-content:active {
  cursor: grabbing;
}

.file-image {
  width: 520px;
  height: 720px;
  object-fit: contain;
  border-radius: 4px;
  transition: transform 0.1s ease-out;
  pointer-events: none;
  user-select: none;
}

.zoom-controls {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 8px;
  background: rgba(0, 0, 0, 0.6);
  padding: 8px;
  border-radius: 8px;
  z-index: 12;
}

.zoom-btn {
  width: 36px;
  height: 36px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.9);
  border: 1px solid rgba(0, 0, 0, 0.1);
  color: #2c3e50;
  font-size: 20px;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  user-select: none;
}

.zoom-btn:hover {
  background: #d50000;
  color: white;
  border-color: #d50000;
  transform: scale(1.1);
}

.zoom-btn:active {
  transform: scale(0.95);
}

.close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.6);
  border: 2px solid rgba(255, 255, 255, 0.3);
  color: white;
  font-size: 28px;
  font-weight: 300;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  z-index: 11;
}

.close-btn:hover {
  background: rgba(213, 0, 0, 0.8);
  border-color: #d50000;
  transform: scale(1.1);
}

.close-btn span {
  line-height: 1;
  margin-top: -2px;
}

/* 轨道颜色主题 */
.track-blue .triangle-logo-large {
  border-bottom-color: #3498db;
  color: #3498db;
}

.track-blue .logo-text,
.track-blue .logo-subtext {
  color: #3498db;
}

.track-yellow .triangle-logo-large {
  border-bottom-color: #f1c40f;
  color: #f1c40f;
}

.track-yellow .logo-text,
.track-yellow .logo-subtext {
  color: #f1c40f;
}

.track-red .triangle-logo-large {
  border-bottom-color: #e74c3c;
  color: #e74c3c;
}

.track-red .logo-text,
.track-red .logo-subtext {
  color: #e74c3c;
}
</style>
