package kubernetes

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	appsv1 "k8s.io/api/apps/v1"
	"sort"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context) ([]*appsv1.Deployment, error)
	Get(ctx context.Context, name string) (*appsv1.Deployment, error)
	Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error)
	Delete(ctx context.Context, name string) error
	Scale(ctx context.Context, name string, replicaNum int32) error
}

type Deployment struct {
	k8s *k8s.Factory
}

var deployment *Deployment

func NewDeployment(k8s *k8s.Factory) *Deployment {
	deployment = &Deployment{k8s: k8s}
	return deployment
}

func (d *Deployment) List(ctx context.Context) ([]*appsv1.Deployment, error) {
	list, err := d.k8s.Deployment().List(ctx)

	if err != nil {
		log.Logger.Error(err)
		return []*appsv1.Deployment{}, fmt.Errorf("record not found")
	}
	//按时间排序
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.GetCreationTimestamp().Time.Before(list[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	return list, nil
}

func (d *Deployment) Get(ctx context.Context, name string) (*appsv1.Deployment, error) {
	dep, err := d.k8s.Deployment().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *Deployment) Create(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
	newDeployment, err := d.k8s.Deployment().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newDeployment, err
}

func (d *Deployment) Update(ctx context.Context, obj *appsv1.Deployment) (*appsv1.Deployment, error) {
	updateDeployment, err := d.k8s.Deployment().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return updateDeployment, err
}

func (d *Deployment) Delete(ctx context.Context, name string) error {
	err := d.k8s.Deployment().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (d *Deployment) Scale(ctx context.Context, name string, replicaNum int32) error {
	err := d.k8s.Deployment().Scale(ctx, name, replicaNum)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}
