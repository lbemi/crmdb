package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/apps/v1"
)

func ListStatefulSets(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).K8S().StatefulSets(namespace).List(c, pageParam, name, label)
}

func GetStatefulSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	statefulSetName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).K8S().StatefulSets(namespace).Get(c, statefulSetName)
}

func GetStatefulSetPods(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	statefulSetName := rc.PathParam("name")
	pods, controllerRevisions := core.V1.Cluster(clusterName).K8S().StatefulSets(namespace).GetStatefulSetPods(c, statefulSetName)
	rc.ResData = map[string]interface{}{
		"pods":                pods,
		"controllerRevisions": controllerRevisions,
	}
}

func GetStatefulSetEvents(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	namespace := rc.PathParam("namespace")
	statefulSetName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).K8S().StatefulSets(namespace).GetStatefulSetEvent(c, statefulSetName)
}

func CreateStatefulSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var statefulSet *v1.StatefulSet
	rc.ShouldBind(&statefulSet)
	rc.ResData = core.V1.Cluster(clusterName).K8S().StatefulSets(statefulSet.Namespace).Create(c, statefulSet)
}

func UpdateStatefulSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var statefulSet *v1.StatefulSet
	rc.ShouldBind(&statefulSet)
	rc.ResData = core.V1.Cluster(clusterName).K8S().StatefulSets(statefulSet.Namespace).Update(c, statefulSet)
}

func DeleteStatefulSet(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	statefulSetName := rc.PathParam("name")
	core.V1.Cluster(clusterName).K8S().StatefulSets(namespace).Delete(c, statefulSetName)
}
