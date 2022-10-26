package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/middleware"
	routes2 "github.com/lbemi/lbemi/routes"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Core())
	defaultRouter := router.Group("/")
	routes2.DefaultRoutes(defaultRouter)

	apiGroup := router.Group("/api/v1", middleware.Test())
	//apiGroup.Use(middleware.JWTAuth(), middleware.CasbinMiddleware())
	apiGroup.Use(middleware.JWTAuth())

	routes2.SetApiGroupRoutes(apiGroup)

	return router
}

func Run() {
	//bootstrap.InitializeConfig()
	//global.App.Log = log.InitializeLog()
	//global.App.Log.Info("log init success!")
	//global.App.Log.Info("监听端口：" + global.App.Config.App.Port)
	//global.App.DB = bootstrap.InitializeDB()
	//global.App.Enforcer = bootstrap.InitCasbinEnforcer()
	//global.App.Redis = bootstrap.InitializeRedis()
	//bootstrap.InitializeValidator()
	//
	//defer func() {
	//	if global.App.DB != nil {
	//		db, _ := global.App.DB.DB()
	//		db.Close()
	//	}
	//}()

	o := option.NewOptions()
	o.Load()

	r := setupRouter()
	//r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	srv := &http.Server{
		Addr:    ":" + o.Config.App.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Logger.Error("listen error", zap.Any("err", err))
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutdown Server ....")
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Logger.Error("Shutdown Server.", zap.Any("error", err))
	}
	log.Logger.Info(" Server exiting....")
}
