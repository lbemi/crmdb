package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/api/cloud"
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
		namespace.GET("", cloud.ListNamespace)
		namespace.GET("/:name", cloud.GetNamespace)
		namespace.POST("", cloud.CreateNamespace)
		namespace.PUT("", cloud.UpdateNamespace)
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
		pod.GET("/log/:namespace/:podName/:container", cloud.GetPodLog)
		pod.GET("/event/:namespace/:name", cloud.GetPodEvents)
	}

	//deployment 资源路由
	deployment := group.Group("/deployment")
	{
		deployment.GET("/:namespace", cloud.ListDeployments)
		deployment.GET("/:namespace/:deploymentName", cloud.GetDeployment)
		deployment.POST("", cloud.CreateDeployment)
		deployment.PUT("", cloud.UpdateDeployment)
		deployment.PUT("/redeploy/:namespace/:name", cloud.ReDeployDeployment)
		deployment.PUT("/rollback/:namespace/:name/:reversion", cloud.RollBackDeployment)
		deployment.DELETE("/:namespace/:deploymentName", cloud.DeleteDeployment)
		deployment.PUT("/:namespace/:deploymentName/:scale", cloud.ScaleDeployments)
		deployment.GET("/:namespace/:deploymentName/pod", cloud.GetDeploymentPods)
		deployment.GET("/:namespace/:deploymentName/event", cloud.GetDeploymentEvents)
	}

	//replicaset 资源路由
	replicaset := group.Group("/replicaset")
	{
		replicaset.GET("/:namespace", cloud.ListReplicaSets)
		replicaset.GET("/:namespace/:name", cloud.GetReplicaSet)
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
		node.PUT("", cloud.UpdateNode)
		node.PATCH("", cloud.PatchNode)
		// 设置是否可以调度
		node.PUT("/:name/:unschedulable", cloud.Schedulable)
		// 排水
		node.POST("/:name/drain", cloud.Drain)
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

	// event 事件路由
	event := group.Group("/event")
	{
		event.GET("/:namespace", cloud.ListEvents)
		event.GET("/:namespace/:name", cloud.GetEvent)
	}

	//  persistentVolumeClaim资源路由
	persistentVolumeClaim := group.Group("/pvc")
	{
		persistentVolumeClaim.GET("/:namespace", cloud.ListPersistentVolumeClaim)
		persistentVolumeClaim.GET("/:namespace/:name", cloud.GetPersistentVolumeClaim)
		persistentVolumeClaim.POST("", cloud.CreatePersistentVolumeClaim)
		persistentVolumeClaim.PUT("", cloud.UpdatePersistentVolumeClaim)
		persistentVolumeClaim.DELETE("/:namespace/:name", cloud.DeletePersistentVolumeClaim)
	}
}
