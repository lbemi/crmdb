package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
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
		name   string            `json:"name"`
		labels map[string]string `json:"labels"`
	}

	var patchData = &patchNode{}
	err := c.ShouldBindJSON(patchData)

	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	list, err := core.V1.Cluster(clusterName).Nodes().Patch(c, patchData.name, patchData.labels)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, list)
}
