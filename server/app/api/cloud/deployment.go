package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
)

func ListDeployments(c *gin.Context) {

	//pageStr := c.DefaultQuery("page", "0")
	//page, err := strconv.Atoi(pageStr)
	//if err != nil {
	//	response.Fail(c, response.ErrCodeParameter)
	//	return
	//}
	//
	//limitStr := c.DefaultQuery("limit", "0")
	//limit, err := strconv.Atoi(limitStr)
	//if err != nil {
	//	response.Fail(c, response.ErrCodeParameter)
	//	return
	//}

	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")
	if namespace == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}
	deploymentList, err := core.V1.Cluster(clusterName).Deployments(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, deploymentList)
}

func GetDeployment(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")
	if namespace == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	deploymentName := c.Param("deploymentName")
	if namespace == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}
	
	deployment, err := core.V1.Cluster(clusterName).Deployments(namespace).Get(c, deploymentName)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, deployment)
}
