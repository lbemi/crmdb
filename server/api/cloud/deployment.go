package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/util"
	v1 "k8s.io/api/apps/v1"
	"strconv"
	"time"
)

func ListDeployments(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

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
	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(deploymentList)

	if pageQuery.Total <= limit {
		pageQuery.Data = deploymentList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = deploymentList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = deploymentList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
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
	//TODO 只能修改某些字段
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var deployment *v1.Deployment

	err := c.ShouldBindJSON(&deployment)
	if err != nil {
		log.Logger.Error(err)
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

func RollBackDeployment(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")
	deploymentName := c.Param("name")
	reversion := c.Param("reversion")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newDeployment, err := core.V1.Cluster(clusterName).Deployments(namespace).RollBack(c, deploymentName, reversion)

	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newDeployment)
}
func ReDeployDeployment(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")
	deploymentName := c.Param("name")

	deployment, err := core.V1.Cluster(clusterName).Deployments(namespace).Get(c, deploymentName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	deployment.Spec.Template.Annotations = map[string]string{
		"lbemi.io/restartAt": time.Now().String(),
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

func ScaleDeployments(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")
	deploymentName := c.Param("deploymentName")
	scale := c.Param("scale")
	scaleNum, err := strconv.Atoi(scale)
	util.GinError(c, err, response.ErrCodeParameter)
	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err = core.V1.Cluster(clusterName).Deployments(namespace).Scale(c, deploymentName, int32(scaleNum))
	util.GinError(c, err, response.ErrCodeParameter)

	response.Success(c, response.StatusOK, nil)
}

func GetDeploymentPods(c *gin.Context) {
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

	pods, replicaSets, err := core.V1.Cluster(clusterName).Deployments(namespace).GetDeploymentPods(c, deploymentName)
	var detail = map[string]interface{}{
		"pods":        pods,
		"replicaSets": replicaSets,
	}
	util.GinError(c, err, response.ErrCodeParameter)
	response.Success(c, response.StatusOK, detail)
}

func GetDeploymentEvents(c *gin.Context) {
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

	events, err := core.V1.Cluster(clusterName).Deployments(namespace).GetDeploymentEvent(c, deploymentName)
	util.GinError(c, err, response.ErrCodeParameter)

	response.Success(c, response.StatusOK, events)
}

func SearchDeployments(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	key := c.DefaultQuery("key", "")
	searchTypeStr := c.DefaultQuery("type", "0")
	searchType, err := strconv.Atoi(searchTypeStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}
	deploymentList, err := core.V1.Cluster(clusterName).Deployments(namespace).Search(c, key, searchType)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(deploymentList)

	if pageQuery.Total <= limit {
		pageQuery.Data = deploymentList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = deploymentList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = deploymentList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}
