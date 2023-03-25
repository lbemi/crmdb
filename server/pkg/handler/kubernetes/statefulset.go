package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/apps/v1"
)

type StatefulSetGetter interface {
	StatefulSets(namespace string) IStatefulSet
}

type IStatefulSet interface {
	List(ctx context.Context) ([]*v1.StatefulSet, error)
	Get(ctx context.Context, name string) (*v1.StatefulSet, error)
	Create(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error)
	Update(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error)
	Delete(ctx context.Context, name string) error
}

type statefulSet struct {
	k8s *k8s.Factory
}

func NewStatefulSet(k8s *k8s.Factory) *statefulSet {
	return &statefulSet{k8s: k8s}
}

func (d *statefulSet) List(ctx context.Context) ([]*v1.StatefulSet, error) {
	list, err := d.k8s.StatefulSet().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}

	return list, err
}

func (d *statefulSet) Get(ctx context.Context, name string) (*v1.StatefulSet, error) {
	sts, err := d.k8s.StatefulSet().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return sts, err
}

func (d *statefulSet) Create(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error) {
	newStatefulSet, err := d.k8s.StatefulSet().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newStatefulSet, err
}

func (d *statefulSet) Update(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error) {
	updateStatefulSet, err := d.k8s.StatefulSet().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return updateStatefulSet, err
}

func (d *statefulSet) Delete(ctx context.Context, name string) error {
	err := d.k8s.StatefulSet().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}
