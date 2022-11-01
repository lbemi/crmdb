import Api from "../api";

//登录
// export const adminLoginApi = (data: LoginReq): Promise<Response> => request.post('/login', data)

// export const getUserLeftMenusApi = (): Promise<Response> => request.get('/user/menus')

// export const getUserListApi = (): Promise<Response> => request.get('/user')

// export const getUserCaptchaApi = (): Promise<Response> => request.get('/captcha')

// export const getUserPermissionApi = (): Promise<Response> => request.get('user/permissions')


export const userApi = {
  login: Api.create('/login','post'),
  listMenus: Api.create('/user/menus','get'),
  listUser: Api.create('/user','get'),
  captcha: Api.create('/captcha', 'get'),
  permission: Api.create('/user/permissions','get')
}