import request from './request'

interface AdminLoginData {
  name: string
  password: string
}
interface Response {
  code: number
  data?: any 
  message: string
}

//登录
export const adminLoginApi = (data:AdminLoginData):Promise<Response>=> request.post('/users/login',data)

export const getUserLeftMenusApi = ():Promise<Response> => request.get('/users/menus')

export const getUserListApi = ():Promise<Response>=> request.get('/users')

