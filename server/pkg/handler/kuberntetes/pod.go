package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type PodGetter interface {
	Pods(namespace string) IPod
}

type IPod interface {
	List(ctx context.Context) ([]*v1.Pod, error)
	Get(ctx context.Context, name string) (*v1.Pod, error)
	Create(ctx context.Context, obj *v1.Pod) (*v1.Pod, error)
	Update(ctx context.Context, obj *v1.Pod) (*v1.Pod, error)
	Delete(ctx context.Context, name string) error
}

type pod struct {
	cli *store.Clients
	ns  string
}

func (d *pod) List(ctx context.Context) ([]*v1.Pod, error) {
	list, err := d.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	return list, err
}

func (d *pod) Get(ctx context.Context, name string) (*v1.Pod, error) {
	dep, err := d.cli.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *pod) Create(ctx context.Context, obj *v1.Pod) (*v1.Pod, error) {
	newPod, err := d.cli.ClientSet.CoreV1().Pods(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newPod, err
}

func (d *pod) Update(ctx context.Context, obj *v1.Pod) (*v1.Pod, error) {
	updatePod, err := d.cli.ClientSet.CoreV1().Pods(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updatePod, err
}

func (d *pod) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.CoreV1().Pods(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func NewPod(cli *store.Clients, namespace string) *pod {
	return &pod{cli: cli, ns: namespace}
}
