import axios from 'axios'

const instance = axios.create({
  baseURL: "http://127.0.0.1:8080/",
  timeout: 15000
})


export interface Response {
  code: number
  result: any
  message?: string
}

// 拦截器
instance.interceptors.request.use(config => {
  let token = sessionStorage.getItem("token")
  if (token) {
    config.headers = config.headers || {} // 判断是会否有headers
    config.headers['Authorization'] = token
  }
  return config
}, err =>{
  return Promise.reject(err)
})

instance.interceptors.response.use(result => {
  return result.data

}, err =>{
  return Promise.reject(err)
})

export default instance