package api

import (
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/core/v1"
)

func ListNodes(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Nodes().List(c, pageParam, name, label)
}

func GetNode(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	nodeName := rc.PathParam("nodeName")
	rc.ResData = core.V1.Cluster(clusterName).Nodes().Get(c, nodeName)
}

func UpdateNode(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var nodeInfo v1.Node
	rc.ShouldBind(&nodeInfo)
	rc.ResData = core.V1.Cluster(clusterName).Nodes().Update(c, &nodeInfo)
}

func PatchNode(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()

	patchData := &entity.PatchNode{}
	rc.ShouldBind(patchData)
	rc.ResData = core.V1.Cluster(clusterName).Nodes().Patch(c, patchData.Name, patchData.Labels)
}

func Schedulable(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	unschedulableStr := rc.PathParam("unschedulable")
	unschedulable := true
	if unschedulableStr == "true" {
		unschedulable = true
	} else if unschedulableStr == "false" {
		unschedulable = false
	}
	rc.ResData = core.V1.Cluster(clusterName).Nodes().Schedulable(c, name, unschedulable)
}

func Drain(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	name := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Nodes().Drain(c, name)
}

func GetPodByNode(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	nodeName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Nodes().GetPodByNode(c, nodeName, pageParam)
}
