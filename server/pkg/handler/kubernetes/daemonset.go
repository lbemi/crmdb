package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/apps/v1"
)

type DaemonSetGetter interface {
	DaemonSets(namespace string) IDaemonSet
}

type IDaemonSet interface {
	List(ctx context.Context) ([]*v1.DaemonSet, error)
	Get(ctx context.Context, name string) (*v1.DaemonSet, error)
	Create(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error)
	Update(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error)
	Delete(ctx context.Context, name string) error
}

type daemonSet struct {
	k8s *k8s.Factory
}

func NewDaemonSet(k8s *k8s.Factory) *daemonSet {
	return &daemonSet{k8s: k8s}
}

func (d *daemonSet) List(ctx context.Context) ([]*v1.DaemonSet, error) {
	list, err := d.k8s.DaemonSet().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *daemonSet) Get(ctx context.Context, name string) (*v1.DaemonSet, error) {
	dep, err := d.k8s.DaemonSet().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *daemonSet) Create(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	newDaemonSet, err := d.k8s.DaemonSet().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newDaemonSet, err
}

func (d *daemonSet) Update(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	updateDaemonSet, err := d.k8s.DaemonSet().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return updateDaemonSet, err
}

func (d *daemonSet) Delete(ctx context.Context, name string) error {
	err := d.k8s.DaemonSet().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}
