package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/apps/v1"
)

type DaemonSetGetter interface {
	DaemonSets(namespace string) IDaemonSet
}

type IDaemonSet interface {
	List(ctx context.Context) []*v1.DaemonSet
	Get(ctx context.Context, name string) *v1.DaemonSet
	Create(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet
	Update(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet
	Delete(ctx context.Context, name string)
}

type daemonSet struct {
	k8s *k8s.Factory
}

func NewDaemonSet(k8s *k8s.Factory) *daemonSet {
	return &daemonSet{k8s: k8s}
}

func (d *daemonSet) List(ctx context.Context) []*v1.DaemonSet {
	return d.k8s.DaemonSet().List(ctx)
}

func (d *daemonSet) Get(ctx context.Context, name string) *v1.DaemonSet {
	return d.k8s.DaemonSet().Get(ctx, name)
}

func (d *daemonSet) Create(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	return d.k8s.DaemonSet().Create(ctx, obj)
}

func (d *daemonSet) Update(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	return d.k8s.DaemonSet().Update(ctx, obj)
}

func (d *daemonSet) Delete(ctx context.Context, name string) {
	d.k8s.DaemonSet().Delete(ctx, name)
}
