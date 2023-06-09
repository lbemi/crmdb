package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
)

type SecretGetter interface {
	Secrets(namespace string) ISecret
}

type ISecret interface {
	List(ctx context.Context) []*v1.Secret
	Get(ctx context.Context, name string) *v1.Secret
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Secret) *v1.Secret
	Update(ctx context.Context, secret *v1.Secret) *v1.Secret
}

type secret struct {
	k8s *k8s.Factory
}

func NewSecret(k8s *k8s.Factory) *secret {
	return &secret{k8s: k8s}
}

func (s *secret) List(ctx context.Context) []*v1.Secret {
	return s.k8s.Secret().List(ctx)
}

func (s *secret) Get(ctx context.Context, name string) *v1.Secret {
	return s.k8s.Secret().Get(ctx, name)
}

func (s *secret) Delete(ctx context.Context, name string) {
	s.k8s.Secret().Delete(ctx, name)
}

func (s *secret) Create(ctx context.Context, secret *v1.Secret) *v1.Secret {
	return s.k8s.Secret().Create(ctx, secret)
}

func (s *secret) Update(ctx context.Context, secret *v1.Secret) *v1.Secret {
	return s.k8s.Secret().Update(ctx, secret)
}
