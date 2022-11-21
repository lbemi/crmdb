package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/app/api/cloud"
)

func NewResourceRoute(group *gin.RouterGroup) {
	resource := group.Group("/resource")
	{
		resource.GET("/deployment/:namespace", cloud.GetDeploymentList)
	}
}
