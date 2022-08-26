package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/middleware"
)

func Run() {
	bootstrap.InitializeConfig()
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	global.App.Log.Info("监听端口：" + global.App.Config.App.Port)
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(false))
	r.GET("/ping", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"message": "pong",
		//})
		panic("asdasds")
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "health",
		})
		global.App.Log.Info("health*****asdaksjdlaksjdklfs")
	})
	r.Run(":" + global.App.Config.App.Port)

}
