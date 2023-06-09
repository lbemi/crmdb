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
//func ListPersistentVolumeClaim(c *gin.Context) {
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
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//	pvcList, err := core.V1.Cluster(clusterName).PersistentVolumeClaim(namespace).List(c)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//
//	// 处理分页 FIXME 待优化
//	var pageQuery types.PageQuery
//	pageQuery.Total = len(pvcList)
//
//	if limit == 0 && page == 0 {
//		pageQuery.Data = pvcList
//	} else {
//		if pageQuery.Total <= limit {
//			pageQuery.Data = pvcList
//		} else if page*limit >= pageQuery.Total {
//			pageQuery.Data = pvcList[(page-1)*limit : pageQuery.Total]
//		} else {
//			pageQuery.Data = pvcList[(page-1)*limit : page*limit]
//		}
//	}
//
//	response.Success(c, response.StatusOK, pageQuery)
//}
//
//func GetPersistentVolumeClaim(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	namespace := c.Param("namespace")
//
//	pvcName := c.Param("name")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	configMap, err := core.V1.Cluster(clusterName).PersistentVolumeClaim(namespace).Get(c, pvcName)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//
//	response.Success(c, response.StatusOK, configMap)
//}
//
//func CreatePersistentVolumeClaim(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	var pvc *v1.PersistentVolumeClaim
//
//	err := c.ShouldBindJSON(&pvc)
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
//	newConfigMap, err := core.V1.Cluster(clusterName).PersistentVolumeClaim(pvc.Namespace).Create(c, pvc)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, newConfigMap)
//}
//
//func UpdatePersistentVolumeClaim(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	var pvc *v1.PersistentVolumeClaim
//
//	err := c.ShouldBindJSON(&pvc)
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
//	newPvcMap, err := core.V1.Cluster(clusterName).PersistentVolumeClaim(pvc.Namespace).Update(c, pvc)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, newPvcMap)
//}
//
//func DeletePersistentVolumeClaim(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//
//	namespace := c.Param("namespace")
//
//	pvcName := c.Param("name")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	err := core.V1.Cluster(clusterName).PersistentVolumeClaim(namespace).Delete(c, pvcName)
//	if err != nil {
//		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
//		return
//	}
//
//	response.Success(c, response.StatusOK, nil)
//}
