package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"strconv"
)

func ListEvents(c *gin.Context) {
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
	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}
	eventList, err := core.V1.Cluster(clusterName).Events(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(eventList)

	if pageQuery.Total <= limit {
		pageQuery.Data = eventList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = eventList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = eventList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}

func GetEvent(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")
	eventName := c.Param("name")

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	event, err := core.V1.Cluster(clusterName).Events(namespace).Get(c, eventName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, event)
}
