package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	appsv1 "k8s.io/api/apps/v1"
)

func ListDaemonSets(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).K8S().DaemonSets(namespace).List(c, pageParam, name, label)
}

func GetDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	daemonSetName := rc.PathParam("daemonSetName")
	rc.ResData = core.V1.Cluster(clusterName).K8S().DaemonSets(namespace).Get(c, daemonSetName)
}

func GetDaemonSetPods(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	daemonSetName := rc.PathParam("name")
	pods, replicaSets := core.V1.Cluster(clusterName).K8S().DaemonSets(namespace).GetDaemonSetPods(c, daemonSetName)
	rc.ResData = map[string]interface{}{
		"pods":        pods,
		"replicaSets": replicaSets,
	}
}

func GetDaemonSetEvents(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	deploymentName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).K8S().DaemonSets(namespace).GetDaemonSetEvent(c, deploymentName)
}

func CreateDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var daemonSet *appsv1.DaemonSet
	rc.ShouldBind(&daemonSet)
	rc.ResData = core.V1.Cluster(clusterName).K8S().DaemonSets(daemonSet.Namespace).Create(c, daemonSet)
}

func UpdateDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var daemonSet *appsv1.DaemonSet
	rc.ShouldBind(&daemonSet)
	rc.ResData = core.V1.Cluster(clusterName).K8S().DaemonSets(daemonSet.Namespace).Update(c, daemonSet)
}

func DeleteDaemonSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	daemonSetName := rc.PathParam("daemonSetName")
	namespace := rc.PathParam("namespace")
	core.V1.Cluster(clusterName).K8S().DaemonSets(namespace).Delete(c, daemonSetName)
}
