package cmd

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/core/server"
	"github.com/lbemi/lbemi/routes/logsys"
	"github.com/lbemi/lbemi/routes/sys"

	//"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var completedOptions *option.Options

func NewDefaultAppCommand() *cobra.Command {
	var configFile string

	rootCmd := &cobra.Command{
		Use:     "GO-OPS system is for operator.",
		Short:   "GO-OPS system is for operator",
		Example: "go-ops --config config.yaml",
		Version: "1.0.0",
		PreRun: func(cmd *cobra.Command, args []string) {
			//初始化
			completedOptions = option.NewOptions().WithConfig(configFile).WithLog().Complete()
			// 注册handler
			core.Register(completedOptions)

			rctx.UserAfterHandlerInterceptor(middleware.LogHandler)
			rctx.UseBeforeHandlerInterceptor(middleware.JWTAuth)
		},
		Run: run,
	}
	rootCmd.Flags().StringVar(&configFile, "config", "", "Set the GO-OPS startup configuration file path")

	return rootCmd
}

func run(cmd *cobra.Command, args []string) {
	httpSever := server.NewHttpSever(":" + completedOptions.Config.App.Port)
	container := httpSever.Container
	container.Filter(middleware.Cors(container).Filter)
	//注册路由
	httpSever.RegisterRoutes(
		sys.UserRoutes(),
		sys.RoleRoutes(),
		sys.MenuRoutes(),
		logsys.LoginLogRoutes(),
		logsys.OperatorLogRoutes(),
	)

	//注册swagger路由
	registerSwaggerRoute(container)

	httpSever.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := httpSever.Stop(); err != nil {
		log.Logger.Errorf("fault stop server. %s", err)
		os.Exit(-3)
	}
}

func registerSwaggerRoute(container *restful.Container) {
	// 注册swagger路由，必须放到最后,否则swagger无法获取所有的路由信息
	config := restfulspec.Config{
		WebServices:                   container.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}

	container.Add(restfulspec.NewOpenAPIService(config))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "GO-OPS",
			Description: "GO-OPS API",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "lbemi",
					Email: "81576870@qq.com",
					URL:   "https://github.com/lbemi/lbemi",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "http://mit.org",
				},
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users***",
		Description: "Managing users"}}}
}
