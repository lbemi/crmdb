package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context) ([]*v1.Deployment, error)
	Get(ctx context.Context, name string) (*v1.Deployment, error)
	Create(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error)
	Update(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error)
	Delete(ctx context.Context, name string) error
}

type deployment struct {
	cli *store.Clients
	ns  string
}

func NewDeployment(cli *store.Clients, namespace string) *deployment {
	return &deployment{cli: cli, ns: namespace}
}

func (d *deployment) List(ctx context.Context) ([]*v1.Deployment, error) {
	list, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	return list, err
}

func (d *deployment) Get(ctx context.Context, name string) (*v1.Deployment, error) {
	dep, err := d.cli.SharedInformerFactory.Apps().V1().Deployments().Lister().Deployments(d.ns).Get(name)
	//dep, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *deployment) Create(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error) {
	newDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newDeployment, err
}

func (d *deployment) Update(ctx context.Context, obj *v1.Deployment) (*v1.Deployment, error) {
	updateDeployment, err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updateDeployment, err
}

func (d *deployment) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.AppsV1().Deployments(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}
