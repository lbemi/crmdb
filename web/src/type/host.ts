export interface HostInfo {
  id: number
  created_at: string
  updated_at: string
  label: string
  remark: string
  ip: string
  port: number
  username: string
  auth_method: number
  status: number
  enable_ssh: number
}

export interface HostForm {
  label: string
  remark: string
  ip: string
  port: number
  username: string
  auth_method: number
  status: number
  enable_ssh: number
  password?: string
  secret?: string
}
