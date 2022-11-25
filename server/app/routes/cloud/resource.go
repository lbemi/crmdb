package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/app/api/cloud"
)

// NewResourceRoute kubernetes 资源路由
func NewResourceRoute(group *gin.RouterGroup) {
	//namespace 资源路由
	namespace := group.Group("/namespace")
	{
		namespace.GET("", cloud.ListNamespace)
		namespace.GET("/:name", cloud.GetNamespace)
		namespace.POST("", cloud.CreateNamespace)
		namespace.DELETE("/:name", cloud.DeleteNamespace)
	}

	//pod 资源路由
	pod := group.Group("/pod")
	{
		pod.GET("/:namespace", cloud.ListPods)
		pod.GET("/:namespace/:podName", cloud.GetPod)
		pod.POST("", cloud.CreatePod)
		pod.PUT("", cloud.UpdatePod)
		pod.DELETE("/:namespace/:podName", cloud.DeletePod)
	}
	//deployment 资源路由
	deployment := group.Group("/deployment")
	{
		deployment.GET("/:namespace", cloud.ListDeployments)
		deployment.GET("/:namespace/:deploymentName", cloud.GetDeployment)
		deployment.POST("", cloud.CreateDeployment)
		deployment.PUT("", cloud.UpdateDeployment)
		deployment.DELETE("/:namespace/:deploymentName", cloud.DeleteDeployment)
	}

	//statefulSet 资源路由
	statefulSet := group.Group("/statefulset")
	{
		statefulSet.GET("/:namespace", cloud.ListStatefulSets)
		statefulSet.GET("/:namespace/:deploymentName", cloud.GetStatefulSet)
		statefulSet.POST("", cloud.CreateStatefulSet)
		statefulSet.PUT("", cloud.UpdateStatefulSet)
		statefulSet.DELETE("/:namespace/:deploymentName", cloud.DeleteStatefulSet)
	}

	//daemonSet 资源路由
	daemonSet := group.Group("/daemonset")
	{
		daemonSet.GET("/:namespace", cloud.ListDaemonSets)
		daemonSet.GET("/:namespace/:deploymentName", cloud.GetDaemonSet)
		daemonSet.POST("", cloud.CreateDaemonSet)
		daemonSet.PUT("", cloud.UpdateDaemonSet)
		daemonSet.DELETE("/:namespace/:deploymentName", cloud.DeleteDaemonSet)
	}

	// node 资源路由
	node := group.Group("/node")
	{
		node.GET("", cloud.ListNodes)
		node.GET("/:nodeName", cloud.GetNode)
	}

	// node 资源路由
	service := group.Group("/service")
	{
		service.GET("/:namespace", cloud.ListServices)
		service.GET("/:namespace/:serviceName", cloud.GetService)
		service.POST("", cloud.CreateService)
		service.PUT("", cloud.UpdateService)
		service.DELETE("/:namespace/:serviceName", cloud.DeleteService)
	}

	// secret 资源路由
	secret := group.Group("/secret")
	{
		secret.GET("/:namespace", cloud.ListSecrets)
		secret.GET("/:namespace/:serviceName", cloud.GetSecret)
		secret.POST("", cloud.CreateSecret)
		secret.PUT("", cloud.UpdateSecret)
		secret.DELETE("/:namespace/:serviceName", cloud.DeleteSecret)
	}

}
