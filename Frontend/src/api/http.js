import axios from 'axios'

const http = axios.create({
  timeout: 10000,
})

http.interceptors.request.use((config) => {
  const token = localStorage.getItem('ta_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

http.interceptors.response.use(
  (response) => response.data,
  (error) => {
    const msg = error.response?.data?.message || error.message || '请求失败'
    return Promise.reject(new Error(msg))
  },
)

export default http

