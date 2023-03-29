package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
)

type ServiceGetter interface {
	Service(namespace string) IService
}

type IService interface {
	List(ctx context.Context) ([]*v1.Service, error)
	Get(ctx context.Context, name string) (*v1.Service, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Service) (*v1.Service, error)
	Update(ctx context.Context, service *v1.Service) (*v1.Service, error)
}

type service struct {
	k8s *k8s.Factory
}

func NewService(k8s *k8s.Factory) *service {
	return &service{k8s: k8s}
}

func (s *service) List(ctx context.Context) ([]*v1.Service, error) {
	nodeList, err := s.k8s.Service().List(ctx)

	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *service) Get(ctx context.Context, name string) (*v1.Service, error) {
	res, err := s.k8s.Service().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *service) Delete(ctx context.Context, name string) error {
	err := s.k8s.Service().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *service) Create(ctx context.Context, service *v1.Service) (*v1.Service, error) {
	res, err := s.k8s.Service().Create(ctx, service)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *service) Update(ctx context.Context, service *v1.Service) (*v1.Service, error) {
	res, err := s.k8s.Service().Update(ctx, service)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}