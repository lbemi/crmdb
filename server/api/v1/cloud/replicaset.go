package cloud

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/lbemi/lbemi/pkg/common/response"
//	"github.com/lbemi/lbemi/pkg/core"
//	"github.com/lbemi/lbemi/pkg/handler/types"
//	"strconv"
//)
//
//func ListReplicaSets(c *gin.Context) {
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
//	replicasetList, err := core.V1.Cluster(clusterName).Replicaset(namespace).List(c)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//	// 处理分页
//	var pageQuery types.PageQuery
//	pageQuery.Total = len(replicasetList)
//
//	if page == 0 && limit == 0 {
//		pageQuery.Data = replicasetList
//	} else {
//		if pageQuery.Total <= limit {
//			pageQuery.Data = replicasetList
//		} else if page*limit >= pageQuery.Total {
//			pageQuery.Data = replicasetList[(page-1)*limit : pageQuery.Total]
//		} else {
//			pageQuery.Data = replicasetList[(page-1)*limit : page*limit]
//		}
//	}
//
//	response.Success(c, response.StatusOK, pageQuery)
//}
//
//func GetReplicaSet(c *gin.Context) {
//	clusterName := c.Query("cloud")
//	if clusterName == "" {
//		response.Fail(c, response.ErrCodeParameter)
//		return
//	}
//	namespace := c.Param("namespace")
//
//	replicasetName := c.Param("name")
//
//	if !core.V1.Cluster(clusterName).CheckHealth() {
//		response.Fail(c, response.ClusterNoHealth)
//		return
//	}
//
//	replicaset, err := core.V1.Cluster(clusterName).Replicaset(namespace).Get(c, replicasetName)
//	if err != nil {
//		response.Fail(c, response.ErrOperateFailed)
//		return
//	}
//
//	response.Success(c, response.StatusOK, replicaset)
//}
