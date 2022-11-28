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
		statefulSet.GET("/:namespace/:statefulSetName", cloud.GetStatefulSet)
		statefulSet.POST("", cloud.CreateStatefulSet)
		statefulSet.PUT("", cloud.UpdateStatefulSet)
		statefulSet.DELETE("/:namespace/:statefulSetName", cloud.DeleteStatefulSet)
	}

	//daemonSet 资源路由
	daemonSet := group.Group("/daemonset")
	{
		daemonSet.GET("/:namespace", cloud.ListDaemonSets)
		daemonSet.GET("/:namespace/:daemonSetName", cloud.GetDaemonSet)
		daemonSet.POST("", cloud.CreateDaemonSet)
		daemonSet.PUT("", cloud.UpdateDaemonSet)
		daemonSet.DELETE("/:namespace/:daemonSetName", cloud.DeleteDaemonSet)
	}

	//job 资源路由
	job := group.Group("/job")
	{
		job.GET("/:namespace", cloud.ListJobs)
		job.GET("/:namespace/:jobName", cloud.GetJob)
		job.POST("", cloud.CreateJob)
		job.PUT("", cloud.UpdateJob)
		job.DELETE("/:namespace/:jobName", cloud.DeleteJob)
	}

	//cronjob 资源路由
	cronjob := group.Group("/cronjob")
	{
		cronjob.GET("/:namespace", cloud.ListCronJobs)
		cronjob.GET("/:namespace/:cronJobName", cloud.GetCronJob)
		cronjob.POST("", cloud.CreateCronJob)
		cronjob.PUT("", cloud.UpdateCronJob)
		cronjob.DELETE("/:namespace/:cronJobName", cloud.DeleteCronJob)
	}

	// node 资源路由
	node := group.Group("/node")
	{
		node.GET("", cloud.ListNodes)
		node.GET("/:nodeName", cloud.GetNode)
	}

	// service 资源路由
	service := group.Group("/service")
	{
		service.GET("/:namespace", cloud.ListServices)
		service.GET("/:namespace/:serviceName", cloud.GetService)
		service.POST("", cloud.CreateService)
		service.PUT("", cloud.UpdateService)
		service.DELETE("/:namespace/:serviceName", cloud.DeleteService)
	}

	// ingress 资源路由
	ingress := group.Group("/ingress")
	{
		ingress.GET("/:namespace", cloud.ListIngresses)
		ingress.GET("/:namespace/:ingressName", cloud.GetIngress)
		ingress.POST("", cloud.CreateIngress)
		ingress.PUT("", cloud.UpdateIngress)
		ingress.DELETE("/:namespace/:ingressName", cloud.DeleteIngress)
	}

	// secret 资源路由
	secret := group.Group("/secret")
	{
		secret.GET("/:namespace", cloud.ListSecrets)
		secret.GET("/:namespace/:secretName", cloud.GetSecret)
		secret.POST("", cloud.CreateSecret)
		secret.PUT("", cloud.UpdateSecret)
		secret.DELETE("/:namespace/:secretName", cloud.DeleteSecret)
	}

	// configmap 资源路由
	configMap := group.Group("/configmap")
	{
		configMap.GET("/:namespace", cloud.ListConfigMaps)
		configMap.GET("/:namespace/:configMapName", cloud.GetConfigMap)
		configMap.POST("", cloud.CreateConfigMap)
		configMap.PUT("", cloud.UpdateConfigMap)
		configMap.DELETE("/:namespace/:configMapName", cloud.DeleteConfigMap)
	}
}
