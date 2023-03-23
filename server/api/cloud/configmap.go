package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	v1 "k8s.io/api/core/v1"
	"strconv"
)

func ListConfigMaps(c *gin.Context) {
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
	configMapList, err := core.V1.Cluster(clusterName).ConfigMaps(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(configMapList)

	if pageQuery.Total <= limit {
		pageQuery.Data = configMapList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = configMapList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = configMapList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}

func GetConfigMap(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")

	configMapName := c.Param("configMapName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	configMap, err := core.V1.Cluster(clusterName).ConfigMaps(namespace).Get(c, configMapName)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, configMap)
}

func CreateConfigMap(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var configMap *v1.ConfigMap

	err := c.ShouldBindJSON(&configMap)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newConfigMap, err := core.V1.Cluster(clusterName).ConfigMaps(configMap.Namespace).Create(c, configMap)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newConfigMap)
}

func UpdateConfigMap(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var configMap *v1.ConfigMap

	err := c.ShouldBindJSON(&configMap)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newConfigMap, err := core.V1.Cluster(clusterName).ConfigMaps(configMap.Namespace).Update(c, configMap)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newConfigMap)
}

func DeleteConfigMap(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	configMapName := c.Param("configMapName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err := core.V1.Cluster(clusterName).ConfigMaps(namespace).Delete(c, configMapName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)
}
