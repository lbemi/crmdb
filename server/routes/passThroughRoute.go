package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/asset"
	"github.com/lbemi/lbemi/api/cloud"
	"github.com/lbemi/lbemi/api/sys"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func PassThroughRoutes(router *gin.RouterGroup) {
	router.POST("/login", sys.Login)
	router.GET("/captcha", sys.GetCaptcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/host/:id/ws", asset.WsShell)
	router.GET("/ws/:cluster/:type", cloud.Ws)
	router.GET("/ws/send", cloud.WsSendAll)

}
