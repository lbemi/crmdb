package app

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"github.com/lbemi/lbemi/routes"
	"github.com/lbemi/lbemi/routes/asset"
	"github.com/lbemi/lbemi/routes/cloud"
	"github.com/lbemi/lbemi/routes/sys"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/middleware"
)

var rootCmd = &cobra.Command{
	Use:     "GO-OPS system is for operator.",
	Short:   "GO-OPS system is for operator",
	Example: "go-ops --config config.yaml",
	Version: "1.0.0",
	PreRun: func(cmd *cobra.Command, args []string) {
		completedOptions := option.NewOptions().Complete()
		initRouter(completedOptions.GinEngine)
	},
	Run: nil,
}

func Run() {

	completedOptions := option.NewOptions().Complete()

	// 注册handler
	core.Register(completedOptions)

	initRouter(completedOptions.GinEngine)
	//r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	srv := &http.Server{
		Addr:    ":" + completedOptions.Config.App.Port,
		Handler: completedOptions.GinEngine,
	}

	defer func() {
		if err := recover(); err != nil {
			log.Logger.Error(err)
			return
		}
	}()

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Logger.Error("listen error", zap.Any("err", err))
		}
	}()

	log.Logger.Infof("启动成功：http://127.0.0.1/%v", completedOptions.Config.App.Port)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutdown Server ....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Error("Shutdown Server.", zap.Any("error", err))
	}

	log.Logger.Info(" Server exiting....")
}

func initRouter(router *gin.Engine) {
	router.Use(middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Cross())

	v1 := router.Group("/api/v1")

	// 注册不需要鉴权路由
	routes.PassThroughRoutes(v1)
	// 中间件middleware.CasbinMiddleware()
	//v1.Use(middleware.JWTAuth())

	//注册业务路由
	//sys.NewUserRouter(v1)
	sys.NewMenuRouter(v1)
	sys.NewRoleRouter(v1)
	asset.NewHostRouter(v1)

	//注册kubernetes 路由
	cloud.NewClusterRoutes(v1)
	cloud.NewResourceRoute(v1)
}
