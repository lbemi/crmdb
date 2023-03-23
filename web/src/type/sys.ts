export interface UserForm {
  user_name: string
  email: string
  status: number
  password: string
  description: string
  confirmPassword: string
}

export interface PageInfo {
  page: number
  limit: number
}

export interface UserInfo {
  id: number
  created_at: string
  updated_at: string
  user_name: string
  email: string
  status: number
  description: string
}

export interface RoleInfo {
  id: number
  created_at: string
  updated_at: string
  name: string
  status: number
  memo: string
  sequence: number
  parent_id: number
  children: MenuObj[]
}

export interface RoleFrom {
  name: string
  status: number
  memo: string
  sequence: number
  parent_id?: number | string
}

export interface MenuInfo {
  id: number
  gmt_create: string
  gmt_modified: string
  resource_version: number
  code: string
  status: number
  memo: string
  parent_id: number
  url: string
  name: string
  icon: string
  sequence: number
  menu_type: number
  method: string
  children: MenuObj[]
}

export interface MenuFrom {
  code?: string
  status: number
  memo: string
  parent_id?: number | string
  url: string
  name: string
  sequence: number
  menu_type: number
  method?: string
  icon?: string
}
