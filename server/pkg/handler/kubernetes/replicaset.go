package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	appsv1 "k8s.io/api/apps/v1"
)

type ReplicasetGetter interface {
	Replicaset(namespace string) ReplicasetImp
}

type ReplicasetImp interface {
	List(ctx context.Context) []*appsv1.ReplicaSet
	Get(ctx context.Context, name string) *appsv1.ReplicaSet
}

type replicaset struct {
	k8s *k8s.Factory
}

func NewReplicaset(k8s *k8s.Factory) *replicaset {
	return &replicaset{k8s: k8s}
}

func (r *replicaset) List(ctx context.Context) []*appsv1.ReplicaSet {
	return r.k8s.Replicaset().List(ctx)
}

func (r *replicaset) Get(ctx context.Context, name string) *appsv1.ReplicaSet {
	return r.k8s.Replicaset().Get(ctx, name)
}
