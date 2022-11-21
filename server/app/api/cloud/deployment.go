package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
)

func GetDeploymentList(c *gin.Context) {
	namespace := c.Param("namespace")
	if namespace == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	deploymentList, err := core.V1.Cluster("test").Deployments(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, deploymentList)
}
