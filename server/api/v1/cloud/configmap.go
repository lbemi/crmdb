package cloud

import (
	v1 "k8s.io/api/core/v1"

	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func ListConfigMaps(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).ConfigMaps(namespace).List(c, pageParam, name, label)
}

func GetConfigMap(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	namespace := rc.PathParam("namespace")
	configMapName := rc.PathParam("configMapName")
	rc.ResData = core.V1.Cluster(clusterName).ConfigMaps(namespace).Get(c, configMapName)
}

func CreateConfigMap(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	var configMap *v1.ConfigMap
	rc.ShouldBind(&configMap)
	rc.ResData = core.V1.Cluster(clusterName).ConfigMaps(configMap.Namespace).Create(c, configMap)
}

func UpdateConfigMap(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	var configMap *v1.ConfigMap
	rc.ShouldBind(&configMap)
	rc.ResData = core.V1.Cluster(clusterName).ConfigMaps(configMap.Namespace).Update(c, configMap)
}

func DeleteConfigMap(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.Query("cloud")
	namespace := rc.PathParam("namespace")
	configMapName := rc.PathParam("name")
	core.V1.Cluster(clusterName).ConfigMaps(namespace).Delete(c, configMapName)
}
