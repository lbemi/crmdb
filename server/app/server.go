package app

import (
	"context"
	"github.com/lbemi/lbemi/app/routes/sys"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lbemi/lbemi/app/option"
	"github.com/lbemi/lbemi/app/routes"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/lbemi"
	"github.com/lbemi/lbemi/pkg/middleware"
)

func Run() {

	o := option.NewOptions()

	o.Load()
	// 注册handler
	lbemi.Setup(o)

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

func initRouter(router *gin.Engine) {
	router.Use(middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Cross())

	v1 := router.Group("/api/v1")
	// 注册默认路由
	routes.DefaultRoutes(v1)
	//, middleware.CasbinMiddleware()
	v1.Use(middleware.JWTAuth(), middleware.CasbinMiddleware())

	//注册业务路由
	sys.NewUserRouter(v1)
	sys.NewMenuRouter(v1)
	sys.NewRoleRouter(v1)

}
