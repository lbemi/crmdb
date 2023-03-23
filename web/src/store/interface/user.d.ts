export interface User {
  id: number
  created_at: string
  updated_at: string
  user_name: string
  email: string
  mobile: string
  status: number
  description: string
}

export interface Data {
  token: string
  user: User
}

export interface LoginResponseObj {
  code: number
  data: Data
  message: string
}

export interface LoginReq {
  user_name: string
  password: string
  captcha: string
  captcha_id: string
}
