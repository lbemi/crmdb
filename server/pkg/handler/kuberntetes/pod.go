package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodGetter interface {
	Pods(namespace string) IPod
}

type IPod interface {
	List(ctx context.Context) (*v1.PodList, error)
	Get(ctx context.Context, name string) (*v1.Pod, error)
	Create(ctx context.Context, obj *v1.Pod) (*v1.Pod, error)
	Update(ctx context.Context, obj *v1.Pod) (*v1.Pod, error)
	Delete(ctx context.Context, name string) error
}

type pod struct {
	cli *cloud.Clients
	ns  string
}

func (d *pod) List(ctx context.Context) (*v1.PodList, error) {
	list, err := d.cli.ClientSet.CoreV1().Pods(d.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}

	return list, err
}

func (d *pod) Get(ctx context.Context, name string) (*v1.Pod, error) {
	dep, err := d.cli.ClientSet.CoreV1().Pods(d.ns).Get(ctx, name, metav1.GetOptions{})
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

func NewPod(cli *cloud.Clients, namespace string) *pod {
	return &pod{cli: cli, ns: namespace}
}
