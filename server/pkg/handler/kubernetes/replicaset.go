package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	appsv1 "k8s.io/api/apps/v1"
)

type ReplicasetGetter interface {
	Replicaset(namespace string) ReplicasetImp
}

type ReplicasetImp interface {
	List(ctx context.Context) ([]*appsv1.ReplicaSet, error)
	Get(ctx context.Context, name string) (*appsv1.ReplicaSet, error)
}

type replicaset struct {
	k8s *k8s.Factory
}

func NewReplicaset(k8s *k8s.Factory) *replicaset {
	return &replicaset{k8s: k8s}
}

func (r *replicaset) List(ctx context.Context) ([]*appsv1.ReplicaSet, error) {
	replicaSets, err := r.k8s.Replicaset().List(ctx)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return replicaSets, nil
}

func (r *replicaset) Get(ctx context.Context, name string) (*appsv1.ReplicaSet, error) {
	replicaSet, err := r.k8s.Replicaset().Get(ctx, name)
	if err != nil {
		log.Logger.Error(replicaSet)
		return nil, err
	}
	return replicaSet, nil
}
