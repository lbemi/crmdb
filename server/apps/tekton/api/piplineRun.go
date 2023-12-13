package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func ListPipelineRuns(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.Query("name")
	label := rc.Query("label")
	pageParam := rc.GetPageQueryParam()
	rc.ResData = core.V1.Cluster(clusterName).Tekton().PipelineRuns(namespace).List(c, pageParam, name, label)
}

func GetPipelineRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	PipelineRunName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Tekton().PipelineRuns(namespace).Get(c, PipelineRunName)
}

func CreatePipelineRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var PipelineRun *v1.PipelineRun
	rc.ShouldBind(&PipelineRun)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().PipelineRuns(PipelineRun.Namespace).Create(c, PipelineRun)
}

func UpdatePipelineRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var PipelineRun *v1.PipelineRun
	rc.ShouldBind(&PipelineRun)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().PipelineRuns(PipelineRun.Namespace).Update(c, PipelineRun)
}

func DeletePipelineRun(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	core.V1.Cluster(clusterName).Tekton().PipelineRuns(namespace).Delete(c, name)
}
