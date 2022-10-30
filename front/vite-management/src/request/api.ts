import request from './request'

interface AdminLoginData {
  user_name: string
  password: string
  captcha: string,
  captcha_id: string
}
interface Response {
  code: number
  data?: any
  message: string
}

//登录
export const adminLoginApi = (data: AdminLoginData): Promise<Response> => request.post('/login', data)

export const getUserLeftMenusApi = (): Promise<Response> => request.get('/user/menus')

export const getUserListApi = (): Promise<Response> => request.get('/user')

export const getUserCaptchaApi = (): Promise<Response> => request.get('/captcha')

export const getUserPermissionApi = (): Promise<Response> => request.get('user/permissions')