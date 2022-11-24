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

	//deployment 资源路由
	deployment := group.Group("/deployment")
	{
		deployment.GET("/:namespace", cloud.ListDeployments)
		deployment.GET("/:namespace/:deploymentName", cloud.GetDeployment)
		deployment.POST("", cloud.CreateDeployment)
		deployment.PUT("", cloud.UpdateDeployment)
		deployment.DELETE("/:namespace/:deploymentName", cloud.DeleteDeployment)
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
