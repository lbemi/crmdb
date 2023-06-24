package asset

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/asset"
)

func NewHostRouter(router *gin.RouterGroup) {
	host := router.Group("/host")
	{
		host.POST("", asset.AddHost)          // 添加主机
		host.GET("", asset.ListHosts)         // 获取主机列表
		host.GET("/:id", asset.GetHostById)   // 根据id获取主机
		host.PUT("/:id", asset.UpdateHost)    // 根据id修改主机
		host.DELETE("/:id", asset.DeleteHost) //根据id删除主机

		//host.GET("/:id/ws", asset.WsShell)
	}
}
