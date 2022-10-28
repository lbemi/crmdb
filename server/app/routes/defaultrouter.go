package routes

import (
	"github.com/gin-gonic/gin"
	sys2 "github.com/lbemi/lbemi/app/api/sys"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DefaultRoutes(router *gin.RouterGroup) {
	router.POST("/login", sys2.Login)

	router.GET("/captcha", sys2.GetCaptcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
