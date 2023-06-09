package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	v1 "k8s.io/api/batch/v1"
	"strconv"
)

func ListCronJobs(c *gin.Context) {
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
	cronJobList, err := core.V1.Cluster(clusterName).CronJobs(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(cronJobList)

	if pageQuery.Total <= limit {
		pageQuery.Data = cronJobList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = cronJobList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = cronJobList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}

func GetCronJob(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")

	cronJobName := c.Param("cronJobName")

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	cronJob, err := core.V1.Cluster(clusterName).CronJobs(namespace).Get(c, cronJobName)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, cronJob)
}

func CreateCronJob(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	var cronJob *v1.CronJob

	err := c.ShouldBindJSON(&cronJob)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newCronJob, err := core.V1.Cluster(clusterName).CronJobs(cronJob.Namespace).Create(c, cronJob)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newCronJob)
}

func UpdateCronJob(c *gin.Context) {
	//TODO 只能修改某些字段
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var cronJob *v1.CronJob

	err := c.ShouldBindJSON(&cronJob)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newCronJob, err := core.V1.Cluster(clusterName).CronJobs(cronJob.Namespace).Update(c, cronJob)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newCronJob)
}

func DeleteCronJob(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	cronJobName := c.Param("cronJobName")

	if !core.V1.Cluster(clusterName).CheckHealth() {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err := core.V1.Cluster(clusterName).CronJobs(namespace).Delete(c, cronJobName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)
}
