package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
)

type ConfigMapGetter interface {
	ConfigMaps(namespace string) IConfigMap
}

type IConfigMap interface {
	List(ctx context.Context) []*v1.ConfigMap
	Get(ctx context.Context, name string) *v1.ConfigMap
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.ConfigMap) *v1.ConfigMap
	Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap
}

type configMap struct {
	k8s *k8s.Factory
}

func NewConfigMap(k8s *k8s.Factory) *configMap {
	return &configMap{k8s: k8s}
}

func (s *configMap) List(ctx context.Context) []*v1.ConfigMap {
	return s.k8s.ConfigMap().List(ctx)
}

func (s *configMap) Get(ctx context.Context, name string) *v1.ConfigMap {
	return s.k8s.ConfigMap().Get(ctx, name)
}

func (s *configMap) Delete(ctx context.Context, name string) {
	s.k8s.ConfigMap().Delete(ctx, name)
}

func (s *configMap) Create(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	return s.k8s.ConfigMap().Create(ctx, configMap)
}

func (s *configMap) Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	return s.k8s.ConfigMap().Update(ctx, configMap)
}
