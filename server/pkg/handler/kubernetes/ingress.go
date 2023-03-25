package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/networking/v1"
)

type IngressesGetter interface {
	Ingresses(namespace string) IIngresses
}

type IIngresses interface {
	List(ctx context.Context) ([]*v1.Ingress, error)
	Get(ctx context.Context, name string) (*v1.Ingress, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Ingress) (*v1.Ingress, error)
	Update(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error)
}

type ingresses struct {
	k8s *k8s.Factory
}

func NewIngresses(k8s *k8s.Factory) *ingresses {
	return &ingresses{k8s: k8s}
}

func (s *ingresses) List(ctx context.Context) ([]*v1.Ingress, error) {
	nodeList, err := s.k8s.Ingress().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *ingresses) Get(ctx context.Context, name string) (*v1.Ingress, error) {
	res, err := s.k8s.Ingress().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *ingresses) Delete(ctx context.Context, name string) error {
	err := s.k8s.Ingress().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *ingresses) Create(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error) {
	res, err := s.k8s.Ingress().Create(ctx, ingresses)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *ingresses) Update(ctx context.Context, ingresses *v1.Ingress) (*v1.Ingress, error) {
	res, err := s.k8s.Ingress().Update(ctx, ingresses)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}
