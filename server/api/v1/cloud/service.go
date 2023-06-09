package cloud

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/lbemi/lbemi/pkg/common/response"
//	"github.com/lbemi/lbemi/pkg/core"
//	"github.com/lbemi/lbemi/pkg/handler/types"
//	v1 "k8s.io/api/core/v1"
//	"strconv"
//)
//
//func ListServices(c *gin.Context) {
//
//	pageStr := c.DefaultQuery("page", "0")
//	page, err := strconv.Atoi(pageStr)
//	if err != nil {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	limitStr := c.DefaultQuery("limit", "0")
//	limit, err := strconv.Atoi(limitStr)
//	if err != nil {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	namespace := c.Param("namespace")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//	serviceList, err := core.V1.Cluster(clusterName).Service(namespace).List(c)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//
//	// 处理分页
//	var pageQuery types.PageQuery
//	pageQuery.Total = len(serviceList)
//
//	if pageQuery.Total <= limit {
//		pageQuery.Data = serviceList
//	} else if page*limit >= pageQuery.Total {
//		pageQuery.Data = serviceList[(page-1)*limit : pageQuery.Total]
//	} else {
//		pageQuery.Data = serviceList[(page-1)*limit : page*limit]
//	}
//
//	response.Success(c, response.StatusOK, pageQuery)
//}
//
//func GetService(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	namespace := c.Param("namespace")
//	serviceName := c.Param("serviceName")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	service, err := core.V1.Cluster(clusterName).Service(namespace).Get(c, serviceName)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, service)
//}
//
//func GetServiceWorkLoad(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	namespace := c.Param("namespace")
//	serviceName := c.Param("serviceName")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	data, err := core.V1.Cluster(clusterName).Service(namespace).ListWorkLoad(c, serviceName)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, data)
//}
//
//func CreateService(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	var service v1.Service
//	err := c.ShouldBindJSON(&service)
//	if err != nil {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	newService, err := core.V1.Cluster(clusterName).Service(service.Namespace).Create(c, &service)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, newService)
//}
//
//func UpdateService(c *gin.Context) {
//	// TODO 不存在service则会创建，优化，不允许修改service名称
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	var service v1.Service
//	err := c.ShouldBindJSON(&service)
//	if err != nil {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	newService, err := core.V1.Cluster(clusterName).Service(service.Namespace).Update(c, &service)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, newService)
//}
//
//func DeleteService(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	namespace := c.Param("namespace")
//	serviceName := c.Param("serviceName")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	err := core.V1.Cluster(clusterName).Service(namespace).Delete(c, serviceName)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, nil)
//}
