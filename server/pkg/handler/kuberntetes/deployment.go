package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentGetter interface {
	Deployments(namespace string) IDeployment
}

type IDeployment interface {
	List(ctx context.Context) (*v1.DeploymentList, error)
	Get(ctx context.Context, name string) (*v1.Deployment, error)
}

type deployment struct {
	client *kubernetes.Clientset
	ns     string
}

func (d *deployment) List(ctx context.Context) (*v1.DeploymentList, error) {
	list, err := d.client.AppsV1().Deployments(d.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *deployment) Get(ctx context.Context, name string) (*v1.Deployment, error) {
	dep, err := d.client.AppsV1().Deployments(d.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func NewDeployment(client *kubernetes.Clientset, namespace string) *deployment {
	return &deployment{client: client, ns: namespace}
}
