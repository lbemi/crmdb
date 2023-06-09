package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/apps/v1"
)

type StatefulSetGetter interface {
	StatefulSets(namespace string) IStatefulSet
}

type IStatefulSet interface {
	List(ctx context.Context) []*v1.StatefulSet
	Get(ctx context.Context, name string) *v1.StatefulSet
	Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Delete(ctx context.Context, name string)
}

type statefulSet struct {
	k8s *k8s.Factory
}

func NewStatefulSet(k8s *k8s.Factory) *statefulSet {
	return &statefulSet{k8s: k8s}
}

func (d *statefulSet) List(ctx context.Context) []*v1.StatefulSet {
	return d.k8s.StatefulSet().List(ctx)
}

func (d *statefulSet) Get(ctx context.Context, name string) *v1.StatefulSet {
	return d.k8s.StatefulSet().Get(ctx, name)
}

func (d *statefulSet) Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	return d.k8s.StatefulSet().Create(ctx, obj)
}

func (d *statefulSet) Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	return d.k8s.StatefulSet().Update(ctx, obj)
}

func (d *statefulSet) Delete(ctx context.Context, name string) {
	d.k8s.StatefulSet().Delete(ctx, name)
}
