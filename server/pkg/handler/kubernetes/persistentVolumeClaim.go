package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaimGetter interface {
	PersistentVolumeClaim(namespace string) PersistentVolumeClaimImp
}

type PersistentVolumeClaimImp interface {
	List(ctx context.Context) ([]*v1.PersistentVolumeClaim, error)
	Get(ctx context.Context, name string) (*v1.PersistentVolumeClaim, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
	Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
}

type persistentVolumeClaim struct {
	k8s *k8s.Factory
}

func NewPersistentVolumeClaim(k8s *k8s.Factory) *persistentVolumeClaim {
	return &persistentVolumeClaim{k8s: k8s}
}

func (p *persistentVolumeClaim) List(ctx context.Context) ([]*v1.PersistentVolumeClaim, error) {
	pvcList, err := p.k8s.PersistentVolumeClaim().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	return pvcList, err
}

func (p *persistentVolumeClaim) Get(ctx context.Context, name string) (*v1.PersistentVolumeClaim, error) {
	res, err := p.k8s.PersistentVolumeClaim().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (p *persistentVolumeClaim) Delete(ctx context.Context, name string) error {
	err := p.k8s.PersistentVolumeClaim().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (p *persistentVolumeClaim) Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	res, err := p.k8s.PersistentVolumeClaim().Create(ctx, pvc)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (p *persistentVolumeClaim) Update(ctx context.Context, configMap *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	res, err := p.k8s.PersistentVolumeClaim().Update(ctx, configMap)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}
