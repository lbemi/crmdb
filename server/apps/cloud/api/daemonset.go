package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/apps/v1"
)

func ListDaemonSets(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).DaemonSets(namespace).List(c, pageParam, name, label)
}

func GetDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	daemonSetName := rc.PathParam("daemonSetName")
	rc.ResData = core.V1.Cluster(clusterName).DaemonSets(namespace).Get(c, daemonSetName)
}

func CreateDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var daemonSet *v1.DaemonSet
	rc.ShouldBind(&daemonSet)
	rc.ResData = core.V1.Cluster(clusterName).DaemonSets(daemonSet.Namespace).Create(c, daemonSet)
}

func UpdateDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var daemonSet *v1.DaemonSet
	rc.ShouldBind(&daemonSet)
	rc.ResData = core.V1.Cluster(clusterName).DaemonSets(daemonSet.Namespace).Update(c, daemonSet)
}

func DeleteDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	daemonSetName := rc.PathParam("daemonSetName")
	namespace := rc.PathParam("namespace")
	core.V1.Cluster(clusterName).DaemonSets(namespace).Delete(c, daemonSetName)
}
