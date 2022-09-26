package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/routes"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Core())
	defaultRouter := router.Group("/")
	routes.DefaultRoutes(defaultRouter)

	apiGroup := router.Group("/api/v1beat", middleware.Test())
	//apiGroup.Use(middleware.JWTAuth(), middleware.CasbinMiddleware())
	apiGroup.Use(middleware.JWTAuth())

	routes.SetApiGroupRoutes(apiGroup)

	return router
}

func Run() {
	bootstrap.InitializeConfig()
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	global.App.Log.Info("监听端口：" + global.App.Config.App.Port)
	global.App.DB = bootstrap.InitializeDB()
	global.App.Enforcer = bootstrap.InitCasbinEnforcer()
	global.App.Redis = bootstrap.InitializeRedis()
	bootstrap.InitializeValidator()

	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	r := setupRouter()
	//r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			global.App.Log.Error("listen error", zap.Any("err", err))
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.App.Log.Info("Shutdown Server ....")
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	if err := srv.Shutdown(context.Background()); err != nil {
		global.App.Log.Error("Shutdown Server.", zap.Any("error", err))
	}
	global.App.Log.Info(" Server exiting....")
}
