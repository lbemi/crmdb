package v1beat1

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/core/v1"
)

func ListServices(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Service(namespace).List(c, pageParam, name, label)
}

func GetService(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	serviceName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Service(namespace).Get(c, serviceName)
}

func GetServiceWorkLoad(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	serviceName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Service(namespace).ListWorkLoad(c, serviceName)
}

func CreateService(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var service v1.Service
	rc.ShouldBind(&service)
	rc.ResData = core.V1.Cluster(clusterName).Service(service.Namespace).Create(c, &service)
}

func UpdateService(rc *rctx.ReqCtx) {
	// TODO 不存在service则会创建，优化，不允许修改service名称
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var service v1.Service
	rc.ShouldBind(&service)
	rc.ResData = core.V1.Cluster(clusterName).Service(service.Namespace).Update(c, &service)
}

func DeleteService(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	serviceName := rc.PathParam("name")
	core.V1.Cluster(clusterName).Service(namespace).Delete(c, serviceName)
}
