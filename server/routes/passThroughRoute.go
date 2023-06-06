package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/v1/asset"
	cloud2 "github.com/lbemi/lbemi/api/v1/cloud"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func PassThroughRoutes(router *gin.RouterGroup) {
	//router.POST("/login", sys2.Login)
	//router.GET("/captcha", sys2.GetCaptcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/host/:id/ws", asset.WsShell)
	router.GET("/ws/:cluster/:type", cloud2.Ws)
	router.GET("/ws/send", cloud2.WsSendAll)
	router.GET("/pod/:namespace/:podName/:container", cloud2.PodExec)
	router.GET("/pod/:namespace/:podName/:container/log", cloud2.GetPodLog)
}
