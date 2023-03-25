package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type ReplicasetImp interface {
	List(ctx context.Context) ([]*appsv1.ReplicaSet, error)
	Get(ctx context.Context, name string) (*appsv1.ReplicaSet, error)
}

type Replicaset struct {
	cli *store.Clients
	ns  string
}

func (r *Replicaset) List(ctx context.Context) ([]*appsv1.ReplicaSet, error) {
	replicaSets, err := r.cli.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(r.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return replicaSets, nil
}

func (r *Replicaset) Get(ctx context.Context, name string) (*appsv1.ReplicaSet, error) {
	replicaSet, err := r.cli.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(r.ns).Get(name)
	if err != nil {
		log.Logger.Error(replicaSet)
		return nil, err
	}
	return replicaSet, nil
}

func newReplicaset(cli *store.Clients, ns string) *Replicaset {
	return &Replicaset{cli: cli, ns: ns}
}

type ReplicasetHandler struct {
	client      *store.Clients
	clusterName string
}

func NewReplicasetHandler(client *store.Clients, clusterName string) *ReplicasetHandler {
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
