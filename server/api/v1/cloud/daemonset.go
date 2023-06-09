package cloud

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/lbemi/lbemi/pkg/common/response"
//	"github.com/lbemi/lbemi/pkg/core"
//	"github.com/lbemi/lbemi/pkg/handler/types"
//	v1 "k8s.io/api/apps/v1"
//	"strconv"
//)
//
//func ListDaemonSets(c *gin.Context) {
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
//	daemonSetList, err := core.V1.Cluster(clusterName).DaemonSets(namespace).List(c)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//	// 处理分页
//	var pageQuery types.PageQuery
//	pageQuery.Total = len(daemonSetList)
//
//	if pageQuery.Total <= limit {
//		pageQuery.Data = daemonSetList
//	} else if page*limit >= pageQuery.Total {
//		pageQuery.Data = daemonSetList[(page-1)*limit : pageQuery.Total]
//	} else {
//		pageQuery.Data = daemonSetList[(page-1)*limit : page*limit]
//	}
//
//	response.Success(c, response.StatusOK, pageQuery)
//}
//
//func GetDaemonSet(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//	namespace := c.Param("namespace")
//
//	daemonSetName := c.Param("daemonSetName")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	daemonSet, err := core.V1.Cluster(clusterName).DaemonSets(namespace).Get(c, daemonSetName)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//
//	response.Success(c, response.StatusOK, daemonSet)
//}
//
//func CreateDaemonSet(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	var daemonSet *v1.DaemonSet
//
//	err := c.ShouldBindJSON(&daemonSet)
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
//	newDaemonSet, err := core.V1.Cluster(clusterName).DaemonSets(daemonSet.Namespace).Create(c, daemonSet)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, newDaemonSet)
//}
//
//func UpdateDaemonSet(c *gin.Context) {
//	//TODO 只能修改某些字段
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	var daemonSet *v1.DaemonSet
//
//	err := c.ShouldBindJSON(&daemonSet)
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
//	newDaemonSet, err := core.V1.Cluster(clusterName).DaemonSets(daemonSet.Namespace).Update(c, daemonSet)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, newDaemonSet)
//}
//
//func DeleteDaemonSet(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	namespace := c.Param("namespace")
//
//	daemonSetName := c.Param("daemonSetName")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	err := core.V1.Cluster(clusterName).DaemonSets(namespace).Delete(c, daemonSetName)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, nil)
//}
