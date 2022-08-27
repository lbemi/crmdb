package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/routes"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)
	return router
}
func Run() {
	bootstrap.InitializeConfig()
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	global.App.Log.Info("监听端口：" + global.App.Config.App.Port)
	bootstrap.InitializeDB()
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	r := setupRouter()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(false))
	//r.GET("/ping", func(c *gin.Context) {
	//	//c.JSON(200, gin.H{
	//	//	"message": "pong",
	//	//})
	//	panic("asdasds")
	//})
	//r.GET("/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"status": "health",
	//	})
	//})
	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	//r.Run(":" + global.App.Config.App.Port)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			global.App.Log.Error("listen error", zap.Any("err", err))
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.App.Log.Info("Shutdown Server ....")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.App.Log.Error("Shutdown Server.", zap.Any("error", err))
	}
	global.App.Log.Info(" Server exiting....")
}
