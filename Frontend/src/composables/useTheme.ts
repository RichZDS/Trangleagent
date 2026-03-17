import { ref, computed, onMounted } from 'vue'

const THEME_KEY = 'ta_theme'

function getStoredTheme(): 'dark' | 'light' {
  return (localStorage.getItem(THEME_KEY) as 'dark' | 'light') || 'dark'
}

function applyTheme(theme: 'dark' | 'light') {
  document.documentElement.setAttribute('data-theme', theme)
}

// 全局共享主题状态
const theme = ref<'dark' | 'light'>(getStoredTheme())

export function useTheme() {
  const isDark = computed(() => theme.value === 'dark')

  const toggleTheme = () => {
    theme.value = theme.value === 'dark' ? 'light' : 'dark'
    localStorage.setItem(THEME_KEY, theme.value)
    applyTheme(theme.value)
  }

  onMounted(() => {
    applyTheme(theme.value)
  })

  return {
    theme,
    isDark,
    toggleTheme,
  }
}
