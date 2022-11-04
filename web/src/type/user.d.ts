export interface UserForm {
  user_name: string;
  email: string;
  mobile: string;
  status: number;
  password: string;
  description: string;
  confirmPassword: string
}


export interface PageInfo {
  page: number
  limit: number
}