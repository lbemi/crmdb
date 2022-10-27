package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/sys"
)

func NewUserRouter(router *gin.RouterGroup) {
	u := router.Group("/user")
	{
		// 用户退出登录
		u.GET("/logout", sys.Logout)
		// 注册
		u.POST("/register", sys.Register)
		// 根据ID获取用户信息
		u.GET("/:id", sys.GetUserInfoById)
		// 获取用户列表
		u.GET("", sys.GetUserList)
		// 删除
		u.DELETE("/:id", sys.DeleteUserByUserId)
		// 更新
		//u.PUT("/:id", app.updateUser)
		//
		//u.GET("/:id", app.getUser)
		//u.GET("", app.listUsers)
		//
		//// 修改密码
		//u.PUT("/change/:id/password", app.changePassword)
		//// 重置密码
		//u.PUT("/reset/:id/password", app.resetPassword)
		//
		//u.GET("/:id/roles", app.getUserRoles)  // 查询当前用户角色
		//u.POST("/:id/roles", app.setUserRoles) // 根据用户id分配角色
		//
		//// 根据菜单ID获取当前用户的菜单的按钮
		//u.GET("/menus/:id/buttons", app.getButtonsByCurrentUser)
		//// 根据用户ID获取用户的菜单
		//u.GET("/menus", app.getLeftMenusByCurrentUser)
	}

}
