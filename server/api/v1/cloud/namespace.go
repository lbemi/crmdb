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

func ListNamespace(c *gin.Context) {
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
	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}
	namespaceList, err := core.V1.Cluster(clusterName).Namespaces().List(c)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}
	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(namespaceList)
	if page != 0 && limit != 0 {
		if pageQuery.Total <= limit {
			pageQuery.Data = namespaceList
		} else if page*limit >= pageQuery.Total {
			pageQuery.Data = namespaceList[(page-1)*limit : pageQuery.Total]
		} else {
			pageQuery.Data = namespaceList[(page-1)*limit : page*limit]
		}
		response.Success(c, response.StatusOK, pageQuery)

	} else {
		pageQuery.Data = namespaceList
		response.Success(c, response.StatusOK, pageQuery)
	}

}

func GetNamespace(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	name := c.Param("name")

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	namespace, err := core.V1.Cluster(clusterName).Namespaces().Get(c, name)
	if err != nil {
		log.Logger.Error(err)
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, namespace)
}

func DeleteNamespace(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	name := c.Param("name")

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err := core.V1.Cluster(clusterName).Namespaces().Delete(c, name)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func CreateNamespace(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	var namespace v1.Namespace
	err := c.ShouldBindJSON(&namespace)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newNamespace, err := core.V1.Cluster(clusterName).Namespaces().Create(c, &namespace)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newNamespace)
}

func UpdateNamespace(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	var namespace v1.Namespace
	err := c.ShouldBindJSON(&namespace)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newNamespace, err := core.V1.Cluster(clusterName).Namespaces().Update(c, &namespace)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newNamespace)
}
