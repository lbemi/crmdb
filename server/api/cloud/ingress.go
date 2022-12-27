package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	v1 "k8s.io/api/networking/v1"
	"strconv"
)

func ListIngresses(c *gin.Context) {
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
	ingressList, err := core.V1.Cluster(clusterName).Ingresses(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(ingressList)

	if pageQuery.Total <= limit {
		pageQuery.Data = ingressList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = ingressList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = ingressList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}

func GetIngress(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")
	ingressName := c.Param("ingressName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	ingress, err := core.V1.Cluster(clusterName).Ingresses(namespace).Get(c, ingressName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, ingress)
}

func CreateIngress(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	var ingress v1.Ingress
	err := c.ShouldBindJSON(&ingress)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newIngress, err := core.V1.Cluster(clusterName).Ingresses(ingress.Namespace).Create(c, &ingress)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newIngress)
}

func UpdateIngress(c *gin.Context) {
	// TODO 不存在ingress则会创建，优化，不允许修改ingress名称
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var ingress v1.Ingress
	err := c.ShouldBindJSON(&ingress)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newIngress, err := core.V1.Cluster(clusterName).Ingresses(ingress.Namespace).Update(c, &ingress)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newIngress)
}

func DeleteIngress(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")
	ingressName := c.Param("ingressName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err := core.V1.Cluster(clusterName).Ingresses(namespace).Delete(c, ingressName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)
}
