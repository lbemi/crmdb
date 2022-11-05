import Api from "@/request/api";

/*
	user.POST("/logout", sys.Logout) // 注册
  user.POST("/register", sys.Register) // 根据ID获取用户信息
  user.GET("/:id", sys.GetUserInfoById)  // 获取用户列表
  user.GET("", sys.GetUserList)  // 删除
  user.DELETE("/:id", sys.DeleteUserByUserId) // 更新
  user.PUT("/:id", sys.UpdateUser)
  user.GET("/:id/roles", sys.GetUserRoles)  // 查询当前用户角色
  user.POST("/:id/roles", sys.SetUserRoles) // 根据用户id分配角色
  user.GET("/permissions", sys.GetButtonsByCurrentUser)   // 根据菜单ID获取当前用户的权限
  user.GET("/menus", sys.GetLeftMenusByCurrentUser)  // 根据用户ID获取用户的菜单
  user.PUT("/:id/status/:status", sys.UpdateUserStatus)  //修改用户状态
*/
export const userApi = {
  login: Api.create('/login', 'post'),
  listMenus: Api.create('/user/menus', 'get'),
  listUser: Api.create('/user', 'get'),
  listUserRole: Api.create('/user/{id}/roles', 'get'),
  captcha: Api.create('/captcha', 'get'),
  permission: Api.create('/user/permissions', 'get'),
  addUser: Api.create('/user/register', 'post'),
  chageStaus: Api.create('/user/{id}/status/{status}', 'put'),
  deleteUser: Api.create('/user/{id}', 'delete'),
  updateUser: Api.create('/user/{id}', 'put'),
  setUserRole: Api.create('/user/{id}/roles', 'post')
}


/*
  role.POST("", sys.AddRole)          // 添加角色
  role.PUT("/:id", sys.UpdateRole)    // 根据角色ID更新角色信息
  role.DELETE("/:id", sys.DeleteRole) // 删除角色
  role.GET("/:id", sys.GetRole)       // 根据角色ID获取角色信息
  role.GET("", sys.ListRoles)         // 获取所有角色

  role.GET("/:id/menus", sys.GetMenusByRole) // 根据角色ID获取角色权限
  role.POST("/:id/menus", sys.SetRoleMenus)  // 根据角色ID设置角色权限
  role.PUT("/:id/status/:status", sys.UpdateRoleStatus)
*/
export const roleApi = {
  list: Api.create('/role', 'get'),
  update: Api.create("/role/{id}",'put'),
  delete: Api.create("/role/{id}",'delete'),
  get: Api.create("/role/{id}",'get'),
  add: Api.create("/role",'post'),

  getRoleMenus: Api.create('/role/{id}/menus','get'),
  setRoleMenus: Api.create('/role/{id}/menus','post'),
  changeStatus: Api.create('/role/{id}/status/{status}','put'),
}

/*
  menu.POST("", sys.AddMenu)
  menu.PUT("/:id", sys.UpdateMenu)
  menu.DELETE("/:id", sys.DeleteMenu)
  menu.GET("/:id", sys.GetMenu)
  menu.GET("", sys.ListMenus)
  menu.PUT("/:id/status/:status", sys.UpdateMenuStatus)
*/
export const menuApi = {
  list: Api.create('/menu', 'get'),
  update: Api.create("/menu/{id}",'put'),
  delete: Api.create("/menu/{id}",'delete'),
  get: Api.create("/menu/{id}",'get'),
  add: Api.create("/menu",'post'),
  changeStatus: Api.create('/menu/{id}/status/{status}','put'),
}