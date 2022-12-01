import axios, { AxiosRequestConfig, AxiosRequestHeaders } from 'axios'
import { ElMessage } from 'element-plus'
import { ResultEnum } from './enums'
import Api from './api'
import { ErrorCodes, ref } from 'vue'

const service = axios.create({
  baseURL: 'http://127.0.0.1:8080/api/v1',
  // baseURL: "http://127.0.0.1:8090/",
  timeout: 15000
})
export interface Response {
  code: number
  data?: any
  message: string
}

// request请求中间件拦截器
service.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers = config.headers || {} // 判断是会否有headers
      config.headers['Authorization'] = token
    }
    return config
  },
  (err) => {
    return Promise.reject(err)
  }
)

// response返回中间件拦截器
service.interceptors.response.use(
  (result) => {
    const data: Response = result.data
    if (data.code == ResultEnum.SUCCESS) {
      return data
    } else {
      return Promise.reject(data)
    }
  },
  (e: any) => {
    if (e.message) {
      // 对响应错误做点什么
      if (e.message.indexOf('timeout') != -1) {
        ElMessage.error('网络超时')
      } else if (e.message == 'Network Error') {
        ElMessage.error('网络连接错误')
      } else if (e.message.indexOf('404')) {
        ElMessage.error('请求接口找不到')
      } else {
        if (e.response.data) ElMessage.error(e.response.data.message)
        else ElMessage.error('接口路径找不到')
      }
    }
    return Promise.reject(e)
  }
)
export const request = async (
  method: string,
  url: string,
  params: any = null,
  headers: any = null,
  options: any = null
): Promise<any> => {
  if (!url) {
    throw new Error('请求url不能为空')
  }
  const flag = ref(false)
  // 简单判断该url是否是restful风格
  if (url.indexOf('{') != -1) {
    flag.value = true
    url = templateResolve(url, params)
  }

  const query: AxiosRequestConfig = {
    method,
    url
  }
  if (headers) {
    query.headers = headers
  }
  const lowMehtod = method.toLowerCase()

  if (lowMehtod === 'post' || lowMehtod === 'put') {
    if (flag.value) {
      query.data = options
      query.params = params
      flag.value = false
    } else {
      query.data = params
    }
  } else {
    if (!flag.value) {
      query.params = params
    } else {
      query.params = params
    }
  }

  try {
    const res = await service.request(query)
    return res
  } catch (e: any) {
    if (e.message) {
      ElMessage.error(e.message)
    }
    return await Promise.reject(e)
  }
}

export const send = (api: Api, params: any, options: any): Promise<any> => {
  return request(api.method, api.url, params, null, options)
}

export const sendWithHeaders = (
  api: Api,
  params: any,
  headers: any
): Promise<any> => {
  return request(api.method, api.url, params, headers, null)
}

const templateResolve = (template: string, param: any) => {
  return template.replace(/\{\w+\}/g, (word) => {
    const key = word.substring(1, word.length - 1)
    const value = param[key]
    delete param[key]
    if (value != null || value != undefined) {
      return value
    }
    return ''
  })
}

export default { request, send, sendWithHeaders }
