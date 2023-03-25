export interface MenuObj {
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
  sequence: number
  menu_type: number
  method: string
  children: MenuObj[]
}
