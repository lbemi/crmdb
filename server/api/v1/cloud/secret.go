package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	v1 "k8s.io/api/core/v1"
	"strconv"
)

func ListSecrets(c *gin.Context) {

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
	secretList, err := core.V1.Cluster(clusterName).Secrets(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(secretList)
	if limit == 0 && page == 0 {
		pageQuery.Data = secretList
	} else {
		if pageQuery.Total <= limit {
			pageQuery.Data = secretList
		} else if page*limit >= pageQuery.Total {
			pageQuery.Data = secretList[(page-1)*limit : pageQuery.Total]
		} else {
			pageQuery.Data = secretList[(page-1)*limit : page*limit]
		}
	}

	response.Success(c, response.StatusOK, pageQuery)
}

func GetSecret(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")

	secretName := c.Param("secretName")

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	secret, err := core.V1.Cluster(clusterName).Secrets(namespace).Get(c, secretName)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, secret)
}

func CreateSecret(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var secret *v1.Secret

	err := c.ShouldBindJSON(&secret)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newSecret, err := core.V1.Cluster(clusterName).Secrets(secret.Namespace).Create(c, secret)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newSecret)
}

func UpdateSecret(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var secret *v1.Secret

	err := c.ShouldBindJSON(&secret)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newSecret, err := core.V1.Cluster(clusterName).Secrets(secret.Namespace).Update(c, secret)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newSecret)
}

func DeleteSecret(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	secretName := c.Param("secretName")

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err := core.V1.Cluster(clusterName).Secrets(namespace).Delete(c, secretName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)
}
