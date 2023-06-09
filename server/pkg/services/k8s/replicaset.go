package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/restfulx"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type ReplicasetImp interface {
	List(ctx context.Context) []*appsv1.ReplicaSet
	Get(ctx context.Context, name string) *appsv1.ReplicaSet
}

type Replicaset struct {
	cli *store.ClientConfig
	ns  string
}

func (r *Replicaset) List(ctx context.Context) []*appsv1.ReplicaSet {
	replicaSets, err := r.cli.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(r.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return replicaSets
}

func (r *Replicaset) Get(ctx context.Context, name string) *appsv1.ReplicaSet {
	replicaSet, err := r.cli.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(r.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return replicaSet
}

func newReplicaset(cli *store.ClientConfig, ns string) *Replicaset {
	return &Replicaset{cli: cli, ns: ns}
}

type ReplicasetHandler struct {
	client      *store.ClientConfig
	clusterName string
}

func NewReplicasetHandler(client *store.ClientConfig, clusterName string) *ReplicasetHandler {
	return &ReplicasetHandler{client: client, clusterName: clusterName}
}

func (r *ReplicasetHandler) OnAdd(obj interface{}) {
	r.notifyReplicaset(obj)
}

func (r *ReplicasetHandler) OnUpdate(oldObj, newObj interface{}) {
	r.notifyReplicaset(newObj)
}

func (r *ReplicasetHandler) OnDelete(obj interface{}) {
	r.notifyReplicaset(obj)
}

func (r *ReplicasetHandler) notifyReplicaset(obj interface{}) {
	namespace := obj.(*appsv1.ReplicaSet).Namespace
	replicates, err := r.client.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(namespace).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	//按时间排序
	sort.Slice(replicates, func(i, j int) bool {
		return replicates[j].ObjectMeta.GetCreationTimestamp().Time.Before(replicates[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	//fmt.Println(r.clusterName, "-----这个空间-----发生数据变化------------")
	go wsstore.WsClientMap.SendClusterResource(r.clusterName, "replicaset", map[string]interface{}{
		"cluster": r.clusterName,
		"type":    "replicaset",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      replicates,
		},
	})
}
