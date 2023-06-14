package cloud

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	v1 "k8s.io/api/core/v1"
)

func ListPersistentVolumeClaim(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).PersistentVolumeClaim(namespace).List(c, pageParam, name, label)
}

func GetPersistentVolumeClaim(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pvcName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).PersistentVolumeClaim(namespace).Get(c, pvcName)
}

func CreatePersistentVolumeClaim(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var pvc *v1.PersistentVolumeClaim
	rc.ShouldBind(&pvc)
	rc.ResData = core.V1.Cluster(clusterName).PersistentVolumeClaim(pvc.Namespace).Create(c, pvc)
}

func UpdatePersistentVolumeClaim(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var pvc *v1.PersistentVolumeClaim
	rc.ShouldBind(&pvc)
	rc.ResData = core.V1.Cluster(clusterName).PersistentVolumeClaim(pvc.Namespace).Update(c, pvc)
}

func DeletePersistentVolumeClaim(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pvcName := rc.PathParam("name")
	core.V1.Cluster(clusterName).PersistentVolumeClaim(namespace).Delete(c, pvcName)
}
