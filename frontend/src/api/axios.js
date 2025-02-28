import axios from 'axios'

const instance = axios.create({
    baseURL: import.meta.env.VITE_API_BASE,
    timeout: 5000
})

// 请求拦截器
instance.interceptors.request.use(config => {
    const token = localStorage.getItem('jwt')
    if (token) {
        config.headers.Authorization = `Bearer ${token}`
    }
    return config
})

// 响应拦截器
instance.interceptors.response.use(
    response => response.data,
    error => {
        if (error.response?.status === 401) {
            localStorage.removeItem('jwt')
            window.location.href = '/login'
        }
        return Promise.reject(error)
    }
)

export default instance