package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/controller/app"
)

func DefaultRoutes(router *gin.RouterGroup) {
	router.POST("/login", app.Login)
	router.GET("/logout", app.Logout)
}
