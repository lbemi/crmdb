package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/controller/app"
	"net/http"
)

func setApiUserGroupRoutes(router *gin.RouterGroup) {
	router.GET("/logout", app.Logout)
	router.POST("/register", app.Register)
	router.GET("/:id", app.GetUserInfoById)
	router.GET("/info", app.GetUserInfos)
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
}
