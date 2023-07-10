package cmd

import (
	"os"
	"os/signal"
	"syscall"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/core/server"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/routes/asset"
	"github.com/lbemi/lbemi/routes/cloud"
	"github.com/lbemi/lbemi/routes/logsys"
	"github.com/lbemi/lbemi/routes/sys"
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
	registerRoute(httpSever)

	httpSever.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	if err := httpSever.Stop(); err != nil {
		log.Logger.Errorf("fault stop server. %s", err)
		os.Exit(-3)
	}

	//GC分析  go tool trace trace.out  或者 GODEBUG=gctrace=1 go run cmd/main.go
	// f, err := os.Create("/Users/lei/Documents/GitHub/lbemi/server/trace.out")
	// if err != nil {
	// 	log.Logger.Error(err)
	// }
	// defer func(f *os.File) {
	// 	err := f.Close()
	// 	if err != nil {

	// 	}
	// }(f)
	// err = trace.Start(f)
	// if err != nil {
	// 	return
	// }
	// defer trace.Stop()
}

// registerRoute registers the routes for the HTTP server.
//
// It takes a pointer to a `server.HttpSever` as its parameter.
// There is no return value.
func registerRoute(httpSever *server.HttpSever) {
	//注册路由
	httpSever.RegisterRoutes(
		sys.UserRoutes(),
		sys.RoleRoutes(),
		sys.MenuRoutes(),
		asset.GroupRoutes(),
		logsys.LoginLogRoutes(),
		logsys.OperatorLogRoutes(),
		cloud.WebSocketRoutes(),
		asset.HostRotes(),
		asset.ResourceAccountRoutes(),
		asset.AccountRoutes(),
		//k8s集群
		cloud.ClusterRoutes(),
		cloud.KubernetesConfigMapRoutes(),
		cloud.KubernetesCronJobRoutes(),
		cloud.KubernetesDaemonSetRoutes(),
		cloud.KubernetesDeploymentRoutes(),
		cloud.KubernetesEventRoutes(),
		cloud.KubernetesIngressRoutes(),
		cloud.KubernetesJobRoutes(),
		cloud.KubernetesNamespaceRoutes(),
		cloud.KubernetesNodeRoutes(),
		cloud.KubernetesPersistentVolumeClaimRoutes(),
		cloud.KubernetesPodRoutes(),
		cloud.KubernetesReplicaSetRoutes(),
		cloud.KubernetesSecretRoutes(),
		cloud.KubernetesServiceRoutes(),
		cloud.KubernetesStatefulSetRoutes(),
	)

	// 注册swagger路由，必须放到最后,否则swagger无法获取所有的路由信息
	config := restfulspec.Config{
		WebServices:                   httpSever.Container.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}

	httpSever.Container.Add(restfulspec.NewOpenAPIService(config))
}

// enrichSwaggerObject modifies the given Swagger object by adding additional information to it.
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
	//swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
	//	Name:        "users***",
	//	Description: "Managing users"}}}
}
