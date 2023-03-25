package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
)

type SecretGetter interface {
	Secrets(namespace string) ISecret
}

type ISecret interface {
	List(ctx context.Context) ([]*v1.Secret, error)
	Get(ctx context.Context, name string) (*v1.Secret, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, node *v1.Secret) (*v1.Secret, error)
	Update(ctx context.Context, secret *v1.Secret) (*v1.Secret, error)
}

type secret struct {
	k8s *k8s.Factory
}

func NewSecret(k8s *k8s.Factory) *secret {
	return &secret{k8s: k8s}
}

func (s *secret) List(ctx context.Context) ([]*v1.Secret, error) {
	nodeList, err := s.k8s.Secret().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *secret) Get(ctx context.Context, name string) (*v1.Secret, error) {
	res, err := s.k8s.Secret().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *secret) Delete(ctx context.Context, name string) error {
	err := s.k8s.Secret().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *secret) Create(ctx context.Context, secret *v1.Secret) (*v1.Secret, error) {
	res, err := s.k8s.Secret().Create(ctx, secret)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *secret) Update(ctx context.Context, secret *v1.Secret) (*v1.Secret, error) {
	res, err := s.k8s.Secret().Update(ctx, secret)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}
