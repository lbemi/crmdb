package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/global"
)

func Run() {
	bootstrap.InitializeConfig()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + global.App.Config.App.Port)
}
