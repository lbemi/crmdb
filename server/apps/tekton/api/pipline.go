package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func ListPipelines(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.Query("name")
	label := rc.Query("label")
	pageParam := rc.GetPageQueryParam()
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Pipelines(namespace).List(c, pageParam, name, label)
}

func GetPipeline(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	PipelineName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Pipelines(namespace).Get(c, PipelineName)
}

func CreatePipeline(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var Pipeline *v1.Pipeline
	rc.ShouldBind(&Pipeline)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Pipelines(Pipeline.Namespace).Create(c, Pipeline)
}

func UpdatePipeline(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var Pipeline *v1.Pipeline
	rc.ShouldBind(&Pipeline)
	rc.ResData = core.V1.Cluster(clusterName).Tekton().Pipelines(Pipeline.Namespace).Update(c, Pipeline)
}

func DeletePipeline(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	core.V1.Cluster(clusterName).Tekton().Pipelines(namespace).Delete(c, name)
}
