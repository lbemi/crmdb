package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/lbemi/lbemi/pkg/controller/app"
)

func setApiUserGroupRoutes(router *gin.RouterGroup) {
	// 用户退出登录
	router.GET("/logout", app.Logout)
	// 注册
	router.POST("/register", app.Register)
	// 根据ID获取用户信息
	router.GET("/:id", app.GetUserInfoById)
	// 获取用户列表
	router.GET("/info", app.GetUserList)
	// 删除
	router.DELETE("/:id", app.DeleteUserByUserId)
	// 更新
	//	router.PUT("/:id", app.updateUser)
	//
	//	router.GET("/:id", app.getUser)
	//	router.GET("", app.listUsers)
	//
	//	// 修改密码
	//	router.PUT("/change/:id/password", app.changePassword)
	//	// 重置密码
	//	router.PUT("/reset/:id/password", app.resetPassword)
	//
	//	router.GET("/:id/roles", app.getUserRoles)  // 查询当前用户角色
	//	router.POST("/:id/roles", app.setUserRoles) // 根据用户id分配角色
	//
	//	// 根据菜单ID获取当前用户的菜单的按钮
	//	router.GET("/menus/:id/buttons", app.getButtonsByCurrentUser)
	//	// 根据用户ID获取用户的菜单
	//	router.GET("/menus", app.getLeftMenusByCurrentUser)
}
