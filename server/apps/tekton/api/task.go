package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func ListTasks(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.Query("name")
	label := rc.Query("label")
	pageParam := rc.GetPageQueryParam()
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Tasks(namespace).List(c, pageParam, name, label)
}

func GetTask(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	taskName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Tasks(namespace).Get(c, taskName)
}

func CreateTask(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var task *v1.Task
	rc.ShouldBind(&task)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Tasks(task.Namespace).Create(c, task)
}

func UpdateTask(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var task *v1.Task
	rc.ShouldBind(&task)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Tasks(task.Namespace).Update(c, task)
}

func DeleteTask(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	core.V1.Cluster(clusterName).Tekton().Tasks(namespace).Delete(c, name)
}
