package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/batch/v1"
)

func ListJobs(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).K8S().Jobs(namespace).List(c, pageParam, name, label)
}

func GetJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	jobName := rc.PathParam("jobName")
	rc.ResData = core.V1.Cluster(clusterName).K8S().Jobs(namespace).Get(c, jobName)
}
func GetJobSetPods(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	daemonSetName := rc.PathParam("name")
	pods := core.V1.Cluster(clusterName).K8S().Jobs(namespace).GetJobPods(c, daemonSetName)
	rc.ResData = map[string]interface{}{
		"pods": pods,
	}
}

func CreateJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var job *v1.Job
	rc.ShouldBind(&job)
	rc.ResData = core.V1.Cluster(clusterName).K8S().Jobs(job.Namespace).Create(c, job)
}

func UpdateJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var job *v1.Job
	rc.ShouldBind(&job)
	rc.ResData = core.V1.Cluster(clusterName).K8S().Jobs(job.Namespace).Update(c, job)
}

func DeleteJob(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	jobName := rc.PathParam("jobName")
	core.V1.Cluster(clusterName).K8S().Jobs(namespace).Delete(c, jobName)
}
