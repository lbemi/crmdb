package cloud

import (
	"github.com/gin-gonic/gin"
	cloud2 "github.com/lbemi/lbemi/api/v1/cloud"
)

// NewResourceRoute k8s 资源路由
func NewResourceRoute(group *gin.RouterGroup) {
	// websocket
	//ws := group.Group("/ws")
	//{
	//	ws.GET("", cloud.Ws)
	//	ws.GET("/send", cloud.WsSendAll)
	//}

	//namespace 资源路由
	namespace := group.Group("/namespace")
	{
		namespace.GET("", cloud2.ListNamespace)
		namespace.GET("/:name", cloud2.GetNamespace)
		namespace.POST("", cloud2.CreateNamespace)
		namespace.PUT("", cloud2.UpdateNamespace)
		namespace.DELETE("/:name", cloud2.DeleteNamespace)
	}

	//pod 资源路由
	pod := group.Group("/pod")
	{
		pod.GET("/:namespace", cloud2.ListPods)
		pod.GET("/:namespace/search", cloud2.SearchPods)
		pod.GET("/:namespace/:podName", cloud2.GetPod)
		pod.POST("", cloud2.CreatePod)
		pod.PUT("", cloud2.UpdatePod)
		pod.DELETE("/:namespace/:podName", cloud2.DeletePod)
		pod.GET("/log/:namespace/:podName/:container", cloud2.GetPodLog)
		pod.GET("/event/:namespace/:name", cloud2.GetPodEvents)
	}

	//deployment 资源路由
	deployment := group.Group("/deployment")
	{
		deployment.GET("/:namespace", cloud2.ListDeployments)
		deployment.GET("/:namespace/search", cloud2.SearchDeployments)
		deployment.GET("/:namespace/:deploymentName", cloud2.GetDeployment)
		deployment.POST("", cloud2.CreateDeployment)
		deployment.PUT("", cloud2.UpdateDeployment)
		deployment.PUT("/redeploy/:namespace/:name", cloud2.ReDeployDeployment)
		deployment.PUT("/rollback/:namespace/:name/:reversion", cloud2.RollBackDeployment)
		deployment.DELETE("/:namespace/:deploymentName", cloud2.DeleteDeployment)
		deployment.PUT("/:namespace/:deploymentName/:scale", cloud2.ScaleDeployments)
		deployment.GET("/:namespace/:deploymentName/pod", cloud2.GetDeploymentPods)
		deployment.GET("/:namespace/:deploymentName/event", cloud2.GetDeploymentEvents)
	}

	//replicaset 资源路由
	replicaset := group.Group("/replicaset")
	{
		replicaset.GET("/:namespace", cloud2.ListReplicaSets)
		replicaset.GET("/:namespace/:name", cloud2.GetReplicaSet)
	}

	//statefulSet 资源路由
	statefulSet := group.Group("/statefulset")
	{
		statefulSet.GET("/:namespace", cloud2.ListStatefulSets)
		statefulSet.GET("/:namespace/:statefulSetName", cloud2.GetStatefulSet)
		statefulSet.POST("", cloud2.CreateStatefulSet)
		statefulSet.PUT("", cloud2.UpdateStatefulSet)
		statefulSet.DELETE("/:namespace/:statefulSetName", cloud2.DeleteStatefulSet)
	}

	//daemonSet 资源路由
	daemonSet := group.Group("/daemonset")
	{
		daemonSet.GET("/:namespace", cloud2.ListDaemonSets)
		daemonSet.GET("/:namespace/:daemonSetName", cloud2.GetDaemonSet)
		daemonSet.POST("", cloud2.CreateDaemonSet)
		daemonSet.PUT("", cloud2.UpdateDaemonSet)
		daemonSet.DELETE("/:namespace/:daemonSetName", cloud2.DeleteDaemonSet)
	}

	//job 资源路由
	job := group.Group("/job")
	{
		job.GET("/:namespace", cloud2.ListJobs)
		job.GET("/:namespace/:jobName", cloud2.GetJob)
		job.POST("", cloud2.CreateJob)
		job.PUT("", cloud2.UpdateJob)
		job.DELETE("/:namespace/:jobName", cloud2.DeleteJob)
	}

	//cronjob 资源路由
	cronjob := group.Group("/cronjob")
	{
		cronjob.GET("/:namespace", cloud2.ListCronJobs)
		cronjob.GET("/:namespace/:cronJobName", cloud2.GetCronJob)
		cronjob.POST("", cloud2.CreateCronJob)
		cronjob.PUT("", cloud2.UpdateCronJob)
		cronjob.DELETE("/:namespace/:cronJobName", cloud2.DeleteCronJob)
	}

	// node 资源路由
	node := group.Group("/node")
	{
		node.GET("", cloud2.ListNodes)
		node.GET("/:nodeName", cloud2.GetNode)
		node.PUT("", cloud2.UpdateNode)
		node.PATCH("", cloud2.PatchNode)
		// 设置是否可以调度
		node.PUT("/:name/:unschedulable", cloud2.Schedulable)
		// 排水
		node.POST("/:name/drain", cloud2.Drain)
		node.GET("/pods/:nodeName", cloud2.GetPodByNode)
	}

	// service 资源路由
	service := group.Group("/service")
	{
		service.GET("/:namespace", cloud2.ListServices)
		service.GET("/:namespace/:serviceName", cloud2.GetService)
		service.GET("/:namespace/:serviceName/work", cloud2.GetServiceWorkLoad)
		service.POST("", cloud2.CreateService)
		service.PUT("", cloud2.UpdateService)
		service.DELETE("/:namespace/:serviceName", cloud2.DeleteService)
	}

	// ingress 资源路由
	ingress := group.Group("/ingress")
	{
		ingress.GET("/:namespace", cloud2.ListIngresses)
		ingress.GET("/:namespace/:ingressName", cloud2.GetIngress)
		ingress.POST("", cloud2.CreateIngress)
		ingress.PUT("", cloud2.UpdateIngress)
		ingress.DELETE("/:namespace/:ingressName", cloud2.DeleteIngress)
	}

	// secret 资源路由
	secret := group.Group("/secret")
	{
		secret.GET("/:namespace", cloud2.ListSecrets)
		secret.GET("/:namespace/:secretName", cloud2.GetSecret)
		secret.POST("", cloud2.CreateSecret)
		secret.PUT("", cloud2.UpdateSecret)
		secret.DELETE("/:namespace/:secretName", cloud2.DeleteSecret)
	}

	// configmap 资源路由
	configMap := group.Group("/configmap")
	{
		configMap.GET("/:namespace", cloud2.ListConfigMaps)
		configMap.GET("/:namespace/:configMapName", cloud2.GetConfigMap)
		configMap.POST("", cloud2.CreateConfigMap)
		configMap.PUT("", cloud2.UpdateConfigMap)
		configMap.DELETE("/:namespace/:configMapName", cloud2.DeleteConfigMap)
	}

	// event 事件路由
	event := group.Group("/event")
	{
		event.GET("/:namespace", cloud2.ListEvents)
		event.GET("/:namespace/:name", cloud2.GetEvent)
	}

	//  persistentVolumeClaim资源路由
	persistentVolumeClaim := group.Group("/pvc")
	{
		persistentVolumeClaim.GET("/:namespace", cloud2.ListPersistentVolumeClaim)
		persistentVolumeClaim.GET("/:namespace/:name", cloud2.GetPersistentVolumeClaim)
		persistentVolumeClaim.POST("", cloud2.CreatePersistentVolumeClaim)
		persistentVolumeClaim.PUT("", cloud2.UpdatePersistentVolumeClaim)
		persistentVolumeClaim.DELETE("/:namespace/:name", cloud2.DeletePersistentVolumeClaim)
	}
}
