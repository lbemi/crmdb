package cmd

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"

	tekton "github.com/lbemi/lbemi/apps/tekton/router"
	ws "github.com/lbemi/lbemi/apps/websocket/router"
	"github.com/lbemi/lbemi/pkg/global"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"

	asset "github.com/lbemi/lbemi/apps/asset/router"
	cloud "github.com/lbemi/lbemi/apps/cloud/router"
	istio "github.com/lbemi/lbemi/apps/istio/router"
	k8s "github.com/lbemi/lbemi/apps/kubernetes/router"
	logsys "github.com/lbemi/lbemi/apps/log/router"
	"github.com/lbemi/lbemi/apps/system/router"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/core/server"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/util"
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
			// 注册聚合服务
			core.Register(completedOptions)
			// 初始化已存在的kubernetes集群client
			go loadKubernetes()
			rctx.UseBeforeHandlerInterceptor(middleware.JWTAuth)
			rctx.UserAfterHandlerInterceptor(middleware.LogHandler)
		},
		Run: run,
	}
	rootCmd.Flags().StringVar(&configFile, "config", "", "Set the GO-OPS startup configuration file path")
	return rootCmd
}

func run(cmd *cobra.Command, args []string) {
	httpSever := server.NewHttpSever(":" + completedOptions.Config.App.Port)
	container := httpSever.Container
	// 加载中间件
	container.Filter(middleware.Cors(container).Filter)
	//注册路由
	registerRoute(httpSever)

	//启动pprof
	newProfileHttpServer(":9999")
	// 启动服务
	httpSever.Start()

	//监听停止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	if err := httpSever.Stop(); err != nil {
		global.Logger.Errorf("fault stop server. %s", err)
		os.Exit(-3)
	}

}

// registerRoute registers the routes for the HTTP server.
//
// It takes a pointer to a `server.HttpSever` as its parameter.
// There is no return value.
func registerRoute(httpSever *server.HttpSever) {
	//注册路由
	httpSever.RegisterRoutes(
		router.UserRoutes(),
		router.RoleRoutes(),
		router.MenuRoutes(),
		asset.GroupRoutes(),
		logsys.LoginLogRoutes(),
		logsys.OperatorLogRoutes(),

		asset.HostRotes(),
		asset.ResourceAccountRoutes(),
		asset.AccountRoutes(),
		//k8s集群
		cloud.ClusterRoutes(),
		k8s.KubernetesConfigMapRoutes(),
		k8s.KubernetesCronJobRoutes(),
		k8s.KubernetesDaemonSetRoutes(),
		k8s.KubernetesDeploymentRoutes(),
		k8s.KubernetesEventRoutes(),
		k8s.KubernetesIngressRoutes(),
		k8s.KubernetesJobRoutes(),
		k8s.KubernetesNamespaceRoutes(),
		k8s.KubernetesNodeRoutes(),
		k8s.KubernetesPersistentVolumeRoutes(),
		k8s.KubernetesStorageClassRoutes(),
		k8s.KubernetesPersistentVolumeClaimRoutes(),
		k8s.KubernetesPodRoutes(),
		k8s.KubernetesReplicaSetRoutes(),
		k8s.KubernetesSecretRoutes(),
		k8s.KubernetesServiceRoutes(),
		k8s.KubernetesStatefulSetRoutes(),
		// k8s代理路由
		k8s.KubernetesProxyRoutes(),
		//istio路由
		istio.IstioVirtualServiceRoutes(),
		istio.IstioGatewayRoutes(),

		//tekton路由
		tekton.TektonTasksRoutes(),
		tekton.TektonTaskRunsRoutes(),
		tekton.TektonPipelinesRoutes(),
		tekton.TektonPipelineRunsRoutes(),
	)
	// websocket取消压缩
	//httpSever.Container.EnableContentEncoding(false)
	httpSever.Container.Add(ws.WebSocketRoutes())

	// 注册swagger路由，必须放到最后,否则swagger无法获取所有的路由信息
	config := restfulspec.Config{
		WebServices:                   httpSever.Container.RegisteredWebServices(), // you control what commService are visible
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
			global.Logger.Errorf("%s 集群异常，请检查集群. %v", cluster.Name, err)
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

// 开启pprof
func newProfileHttpServer(addr string) {
	go func() {
		global.Logger.Error(http.ListenAndServe(addr, nil))
	}()
}

func newTrace() {
	//GC分析  go tool trace trace.out  或者 GODEBUG=gctrace=1 go run cmd/main.go
	f, err := os.Create("./trace.out")
	if err != nil {
		global.Logger.Error(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			global.Logger.Error(err)
		}
	}(f)

	err = trace.Start(f)
	println("trace start----------------")
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer trace.Stop()
}
