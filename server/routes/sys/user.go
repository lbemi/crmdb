package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/sys"
)

func NewUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		// 用户退出登录
		user.POST("/logout", sys.Logout)
		// 注册
		user.POST("/register", sys.Register)
		// 根据ID获取用户信息
		user.GET("/:id", sys.GetUserInfoById)
		// 获取用户列表
		user.GET("", sys.GetUserList)
		// 删除
		user.DELETE("/:id", sys.DeleteUserByUserId)
		// 更新
		user.PUT("/:id", sys.UpdateUser)

		user.GET("/:id/roles", sys.GetUserRoles)  // 查询当前用户角色
		user.POST("/:id/roles", sys.SetUserRoles) // 根据用户id分配角色

		// 根据菜单ID获取当前用户的权限
		user.GET("/permissions", sys.GetButtonsByCurrentUser)
		// 根据用户ID获取用户的菜单
		user.GET("/menus", sys.GetLeftMenusByCurrentUser)
		//修改用户状态
		user.PUT("/:id/status/:status", sys.UpdateUserStatus)
	}

}
