package routes

import (
	"github.com/lbemi/lbemi/api/sys"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "health")
	})
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
	r := router.Group("/user")
	setApiUserGroupRoutes(r)
	menu := router.Group("menu")
	{
		menu.GET("/allmenus", sys.GetMenuList)
		menu.GET("/list", sys.GetMenuList)
		menu.GET("/detail", sys.GetMenuList)
		menu.GET("/allmenu", sys.GetMenuList)
		menu.GET("/menubuttonlist", sys.GetMenuList)
		menu.POST("/delete", sys.GetMenuList)
		menu.POST("/update", sys.GetMenuList)
		menu.POST("/create", sys.GetMenuList)
	}

}
