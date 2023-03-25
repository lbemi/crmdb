package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
)

type ConfigMapGetter interface {
	ConfigMaps(namespace string) IConfigMap
}

type IConfigMap interface {
	List(ctx context.Context) ([]*v1.ConfigMap, error)
	Get(ctx context.Context, name string) (*v1.ConfigMap, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.ConfigMap) (*v1.ConfigMap, error)
	Update(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error)
}

type configMap struct {
	k8s *k8s.Factory
}

func NewConfigMap(k8s *k8s.Factory) *configMap {
	return &configMap{k8s: k8s}
}

func (s *configMap) List(ctx context.Context) ([]*v1.ConfigMap, error) {
	nodeList, err := s.k8s.ConfigMap().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *configMap) Get(ctx context.Context, name string) (*v1.ConfigMap, error) {
	res, err := s.k8s.ConfigMap().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *configMap) Delete(ctx context.Context, name string) error {
	err := s.k8s.ConfigMap().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *configMap) Create(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	res, err := s.k8s.ConfigMap().Create(ctx, configMap)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *configMap) Update(ctx context.Context, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	res, err := s.Update(ctx, configMap)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}
