package api

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"

	v1 "k8s.io/api/core/v1"
)

func ListSecrets(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).K8S().Secrets(namespace).List(c, pageParam, name, label)
}

func GetSecret(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	secretName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).K8S().Secrets(namespace).Get(c, secretName)
}

func CreateSecret(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var secret *v1.Secret
	rc.ShouldBind(&secret)
	rc.ResData = core.V1.Cluster(clusterName).K8S().Secrets(secret.Namespace).Create(c, secret)
}

func UpdateSecret(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var secret *v1.Secret
	rc.ShouldBind(&secret)
	rc.ResData = core.V1.Cluster(clusterName).K8S().Secrets(secret.Namespace).Update(c, secret)
}

func DeleteSecret(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	secretName := rc.PathParam("name")
	core.V1.Cluster(clusterName).K8S().Secrets(namespace).Delete(c, secretName)

}
