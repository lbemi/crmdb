package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func ListTaskRuns(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.Query("name")
	label := rc.Query("label")
	pageParam := rc.GetPageQueryParam()
	rc.ResData = core.V1.Cluster(clusterName).Tekton().TaskRuns(namespace).List(c, pageParam, name, label)
}

func GetTaskRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	TaskRunName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Tekton().TaskRuns(namespace).Get(c, TaskRunName)
}

func CreateTaskRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var TaskRun *v1.TaskRun
	rc.ShouldBind(&TaskRun)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().TaskRuns(TaskRun.Namespace).Create(c, TaskRun)
}

func UpdateTaskRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var TaskRun *v1.TaskRun
	rc.ShouldBind(&TaskRun)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().TaskRuns(TaskRun.Namespace).Update(c, TaskRun)
}

func DeleteTaskRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	core.V1.Cluster(clusterName).Tekton().TaskRuns(namespace).Delete(c, name)
}
