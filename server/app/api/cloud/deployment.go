package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	v1 "k8s.io/api/apps/v1"
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

	deploymentName := c.Param("deploymentName")

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

func CreateDeployment(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var deployment *v1.Deployment

	err := c.ShouldBindJSON(&deployment)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newDeployment, err := core.V1.Cluster(clusterName).Deployments(deployment.Namespace).Create(c, deployment)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newDeployment)
}

func UpdateDeployment(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var deployment *v1.Deployment

	err := c.ShouldBindJSON(&deployment)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newDeployment, err := core.V1.Cluster(clusterName).Deployments(deployment.Namespace).Update(c, deployment)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newDeployment)
}

func DeleteDeployment(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	deploymentName := c.Param("deploymentName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err := core.V1.Cluster(clusterName).Deployments(namespace).Delete(c, deploymentName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)
}
