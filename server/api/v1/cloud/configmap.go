package cloud

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	v1 "k8s.io/api/core/v1"
)

func ListConfigMaps(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud("cloud")
	pageParam := rc.GetPageQueryParam()
	name := rc.QueryCloud("name")
	label := rc.QueryCloud("label")
	namespace := rc.PathParam("namespace")
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
	configMapName := rc.PathParam("configMapName")
	core.V1.Cluster(clusterName).ConfigMaps(namespace).Delete(c, configMapName)
}
