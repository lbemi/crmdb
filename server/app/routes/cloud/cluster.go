package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/app/api/cloud"
)

func NewClusterRoutes(router *gin.RouterGroup) {

	cluster := router.Group("/cluster")
	{
		cluster.POST("", cloud.CreateCluster)
	}

}
