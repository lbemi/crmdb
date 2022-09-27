package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lbemi/lbemi/pkg/controller/app"
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
		menu.GET("/allmenus", app.GetMenuList)
		menu.GET("/list", app.GetMenuList)
		menu.GET("/detail", app.GetMenuList)
		menu.GET("/allmenu", app.GetMenuList)
		menu.GET("/menubuttonlist", app.GetMenuList)
		menu.POST("/delete", app.GetMenuList)
		menu.POST("/update", app.GetMenuList)
		menu.POST("/create", app.GetMenuList)
	}

}
