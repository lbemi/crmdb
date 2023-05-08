package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/util"
	v1 "k8s.io/api/core/v1"
	"strconv"
)

func ListNodes(c *gin.Context) {

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

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	nodeList, err := core.V1.Cluster(clusterName).Nodes().List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(nodeList)

	if pageQuery.Total <= limit {
		pageQuery.Data = nodeList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = nodeList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = nodeList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}

func GetNode(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	nodeName := c.Param("nodeName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	list, err := core.V1.Cluster(clusterName).Nodes().Get(c, nodeName)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, list)
}

func UpdateNode(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var nodeInfo v1.Node
	err := c.ShouldBindJSON(&nodeInfo)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}
	list, err := core.V1.Cluster(clusterName).Nodes().Update(c, &nodeInfo)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, list)
}

func PatchNode(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	type patchNode struct {
		Name   string                 `json:"name"`
		Labels map[string]interface{} `json:"labels"`
	}

	var patchData patchNode
	err := c.ShouldBindJSON(&patchData)

	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	list, err := core.V1.Cluster(clusterName).Nodes().Patch(c, patchData.Name, patchData.Labels)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, list)
}

func Schedulable(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	name := c.Param("name")
	unschedulableStr := c.Param("unschedulable")
	var unschedulable bool
	if unschedulableStr == "true" {
		unschedulable = true
	} else if unschedulableStr == "false" {
		unschedulable = false
	} else {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	list, err := core.V1.Cluster(clusterName).Nodes().Schedulable(c, name, unschedulable)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, list)
}

func Drain(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	name := c.Param("name")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	res, err := core.V1.Cluster(clusterName).Nodes().Drain(c, name)
	if err != nil || !res {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, "")
}

func GetPodByNode(c *gin.Context) {

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

	nodeName := c.Param("nodeName")
	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	podList, err := core.V1.Cluster(clusterName).Nodes().GetPodByNode(c, nodeName)
	util.GinError(c, err, response.ErrCodeParameter)
	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(podList.Items)

	if pageQuery.Total <= limit {
		pageQuery.Data = podList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = podList.Items[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = podList.Items[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}
