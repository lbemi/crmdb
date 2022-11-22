package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/app/api/cloud"
)

// NewResourceRoute kubernetes 资源路由
func NewResourceRoute(group *gin.RouterGroup) {
	//deployment 资源路由
	deployment := group.Group("/deployment")
	{
		deployment.GET("/:namespace", cloud.GetDeploymentList)
	}
	// node 资源路由
	node := group.Group("/node")
	{
		node.GET("", cloud.ListNodes)
	}
}
