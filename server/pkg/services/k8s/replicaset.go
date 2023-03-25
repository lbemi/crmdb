package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
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
