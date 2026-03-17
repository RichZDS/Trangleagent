/**
 * WebSocket 连接配置
 * 开发环境通过 Vite 代理转发到后端，生产环境使用同源或配置的地址
 */
export function getWsBaseUrl() {
  const env = import.meta.env
  const wsUrl = env.VITE_WS_URL
  if (wsUrl) {
    return (wsUrl.startsWith('http') ? 'ws' + wsUrl.slice(4) : wsUrl).replace(/\/$/, '')
  }
  const { protocol, hostname, port } = window.location
  const wsProtocol = protocol === 'https:' ? 'wss:' : 'ws:'
  // 开发时 Vite 代理 /ws，使用当前 host
  const isDev = hostname === 'localhost' && (port === '5173' || port === '3000')
  return isDev ? `${wsProtocol}//${hostname}:${port}` : `${wsProtocol}//${hostname}${port ? ':' + port : ''}`
}

export function getChatWsUrl() {
  return `${getWsBaseUrl()}/ws/chat`
}
