package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/sys"
)

func NewRoleRouter(router *gin.RouterGroup) {
	role := router.Group("/role")
	{
		role.POST("", sys.AddRole)          // 添加角色
		role.PUT("/:id", sys.UpdateRole)    // 根据角色ID更新角色信息
		role.DELETE("/:id", sys.DeleteRole) // 删除角色
		role.GET("/:id", sys.GetRole)       // 根据角色ID获取角色信息
		role.GET("", sys.ListRoles)         // 获取所有角色

		role.GET("/:id/menus", sys.GetMenusByRole)            // 根据角色ID获取角色权限
		role.POST("/:id/menus", sys.SetRoleMenus)             // 根据角色ID设置角色权限
		role.PUT("/:id/status/:status", sys.UpdateRoleStatus) // 修改角色状态
	}
}
