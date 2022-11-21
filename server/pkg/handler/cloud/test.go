package cloud

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
)

type ResourceGetter interface {
	Resource() IResource
}

type IResource interface {
	GetDeploymentsByNamespace(ctx context.Context, namespace string) (*v1.DeploymentList, error)
	GetPods(ctx context.Context, namespace string) (*v12.PodList, error)
}

type resource struct {
	client services.IDbFactory
}

func (r *resource) GetDeploymentsByNamespace(ctx context.Context, namespace string) (*v1.DeploymentList, error) {
	deploymentList, err := r.client.Resource().GetDeploymentsByNamespace(ctx, namespace)
	if err != nil {
		return nil, err
	}
	return deploymentList, nil
}

func (r *resource) GetPods(ctx context.Context, namespace string) (*v12.PodList, error) {
	pods, err := r.client.Resource().GetPods(ctx, namespace)
	if err != nil {
		return nil, err
	}
	return pods, nil
}

func NewResource(client services.IDbFactory) *resource {
	return &resource{
		client: client,
	}
}
