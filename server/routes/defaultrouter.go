package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/sys"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DefaultRoutes(router *gin.RouterGroup) {
	router.POST("/login", sys.Login)

	router.GET("/captcha", sys.GetCaptcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
