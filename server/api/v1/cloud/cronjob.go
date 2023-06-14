package cloud

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	v1 "k8s.io/api/batch/v1"
)

func ListCronJobs(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.Query("name")
	label := rc.Query("label")
	pageParam := rc.GetPageQueryParam()
	rc.ResData = core.V1.Cluster(clusterName).CronJobs(namespace).List(c, pageParam, name, label)
}

func GetCronJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	cronJobName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).CronJobs(namespace).Get(c, cronJobName)
}

func CreateCronJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var cronJob *v1.CronJob
	rc.ShouldBind(&cronJob)
	rc.ResData = core.V1.Cluster(clusterName).CronJobs(cronJob.Namespace).Create(c, cronJob)
}

func UpdateCronJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var cronJob *v1.CronJob
	rc.ShouldBind(&cronJob)
	rc.ResData = core.V1.Cluster(clusterName).CronJobs(cronJob.Namespace).Update(c, cronJob)
}

func DeleteCronJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.Query("name")
	core.V1.Cluster(clusterName).CronJobs(namespace).Delete(c, name)
}
