package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/app/api/sys"
)

func NewMenuRouter(router *gin.RouterGroup) {
	menu := router.Group("/menu")
	{
		menu.POST("", sys.AddMenu)
		menu.PUT("/:id", sys.UpdateMenu)
		menu.DELETE("/:id", sys.DeleteMenu)
		menu.GET("/:id", sys.GetMenu)
		menu.GET("", sys.ListMenus)
		menu.PUT("/:id/status/:status", sys.UpdateMenuStatus)
	}
}
