export interface UserForm {
  user_name: string;
  email: string;
  status: number;
  password: string;
  description: string;
  confirmPassword: string
}


export interface PageInfo {
  page: number
  limit: number
}


export interface UserInfo {
	id: number;
	created_at: string;
	updated_at: string;
	user_name: string;
	email: string;
	status: number;
	description: string;
}

