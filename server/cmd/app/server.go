package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lbemi/lbemi/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/routes"
	"github.com/lbemi/lbemi/routes/sys/menu"
	"github.com/lbemi/lbemi/routes/sys/user"
)

func initRouter(router *gin.Engine) {
	router.Use(middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Cross())

	v1 := router.Group("/api/v1", middleware.Test())
	// 注册默认路由
	routes.DefaultRoutes(v1)

	//apiGroup.Use(middleware.JWTAuth(), middleware.CasbinMiddleware())
	v1.Use(middleware.JWTAuth())

	//注册业务路由
	user.NewUserRouter(v1)
	menu.NewMenuRouter(v1)

}

func Run() {

	o := option.NewOptions()

	o.Load()
	// 注册handler
	core.Setup(o)

	initRouter(o.GinEngine)
	//r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	srv := &http.Server{
		Addr:    ":" + o.Config.App.Port,
		Handler: o.GinEngine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Logger.Error("listen error", zap.Any("err", err))
		}
		log.Logger.Info("启动成功：", zap.Any("port", o.Config.App.Port))
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
