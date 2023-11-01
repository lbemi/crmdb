package cmd

import (
	"github.com/lbemi/lbemi/pkg/util"
	"github.com/lbemi/lbemi/routes/istio"
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

			// 初始化已存在的kubernetes集群client
			go loadKubernetes()
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

		//istio路由
		istio.IstioVirtualServiceRoutes(),
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

func loadKubernetes() {
	list := core.V1.Cluster("").List()
	for _, cluster := range *list {
		config := util.Decrypt(cluster.KubeConfig)
		_, _, err := core.V1.Cluster("").GenerateClient(cluster.Name, config)
		if err != nil {
			//如果初始化失败且数据库状态是正常，则修改为异常
			if cluster.Status {
				core.V1.Cluster("").ChangeStatus(cluster.ID, false)
			}
			log.Logger.Errorf("%s 集群异常，请检查集群. %v", cluster.Name, err)
			// TODO 是否手设置手动启动监听  启动informer监听
			//go f.Cluster().StartInformer(cluster.Name)
		} else {
			if !cluster.Status {
				// 初始化成功后，如果当前集群状态异常的则改为正常
				core.V1.Cluster("").ChangeStatus(cluster.ID, true)
			}

		}
	}
}
