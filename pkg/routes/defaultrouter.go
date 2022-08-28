package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/lbemi/lbemi/pkg/controller/app"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DefaultRoutes(router *gin.RouterGroup) {
	router.POST("/login", app.Login)
	router.GET("/logout", app.Logout)
	router.GET("/captcha", app.GetCaptcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
