package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/core/v1"
)

func ListPersistentVolume(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).K8S().PersistentVolume().List(c, pageParam, name, label)
}

func GetPersistentVolume(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pvName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).K8S().PersistentVolume().Get(c, pvName)
}

func CreatePersistentVolume(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var pv *v1.PersistentVolume
	rc.ShouldBind(&pv)
	rc.ResData = core.V1.Cluster(clusterName).K8S().PersistentVolume().Create(c, pv)
}

func DeletePersistentVolume(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pvName := rc.PathParam("name")
	core.V1.Cluster(clusterName).K8S().PersistentVolume().Delete(c, pvName)
}
