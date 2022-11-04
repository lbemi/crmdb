import Api from "@/request/api";

export const userApi = {
  login: Api.create('/login','post'),
  listMenus: Api.create('/user/menus','get'),
  listUser: Api.create('/user','get'),
  captcha: Api.create('/captcha', 'get'),
  permission: Api.create('/user/permissions','get'),
  addUser: Api.create('/user/register','post'),
  chageStaus: Api.create('/user/{id}/status/{status}','put'),
  deleteUser: Api.create('/user/{id}','delete'),
}
