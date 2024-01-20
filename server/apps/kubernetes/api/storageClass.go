package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	"k8s.io/api/storage/v1"
)

func ListStorageClass(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).K8S().StorageClass().List(c, pageParam, name, label)
}

func GetStorageClass(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pvName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).K8S().StorageClass().Get(c, pvName)
}

func CreateStorageClass(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var pv *v1.StorageClass
	rc.ShouldBind(&pv)
	rc.ResData = core.V1.Cluster(clusterName).K8S().StorageClass().Create(c, pv)
}

func DeleteStorageClass(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	pvName := rc.PathParam("name")
	core.V1.Cluster(clusterName).K8S().StorageClass().Delete(c, pvName)
}
