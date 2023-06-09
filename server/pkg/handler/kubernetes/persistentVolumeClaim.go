package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaimGetter interface {
	PersistentVolumeClaim(namespace string) PersistentVolumeClaimImp
}

type PersistentVolumeClaimImp interface {
	List(ctx context.Context) []*v1.PersistentVolumeClaim
	Get(ctx context.Context, name string) *v1.PersistentVolumeClaim
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
	Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
}

type persistentVolumeClaim struct {
	k8s *k8s.Factory
}

func NewPersistentVolumeClaim(k8s *k8s.Factory) *persistentVolumeClaim {
	return &persistentVolumeClaim{k8s: k8s}
}

func (p *persistentVolumeClaim) List(ctx context.Context) []*v1.PersistentVolumeClaim {
	return p.k8s.PersistentVolumeClaim().List(ctx)
}

func (p *persistentVolumeClaim) Get(ctx context.Context, name string) *v1.PersistentVolumeClaim {
	return p.k8s.PersistentVolumeClaim().Get(ctx, name)
}

func (p *persistentVolumeClaim) Delete(ctx context.Context, name string) {
	p.k8s.PersistentVolumeClaim().Delete(ctx, name)
}

func (p *persistentVolumeClaim) Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	return p.k8s.PersistentVolumeClaim().Create(ctx, pvc)
}

func (p *persistentVolumeClaim) Update(ctx context.Context, configMap *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	return p.k8s.PersistentVolumeClaim().Update(ctx, configMap)
}
