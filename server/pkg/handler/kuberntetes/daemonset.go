package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSetGetter interface {
	DaemonSets(namespace string) IDaemonSet
}

type IDaemonSet interface {
	List(ctx context.Context) (*v1.DaemonSetList, error)
	Get(ctx context.Context, name string) (*v1.DaemonSet, error)
	Create(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error)
	Update(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error)
	Delete(ctx context.Context, name string) error
}

type daemonSet struct {
	cli *cloud.Clients
	ns  string
}

func (d *daemonSet) List(ctx context.Context) (*v1.DaemonSetList, error) {
	list, err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *daemonSet) Get(ctx context.Context, name string) (*v1.DaemonSet, error) {
	dep, err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *daemonSet) Create(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	newDaemonSet, err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newDaemonSet, err
}

func (d *daemonSet) Update(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	updateDaemonSet, err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updateDaemonSet, err
}

func (d *daemonSet) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func NewDaemonSet(cli *cloud.Clients, namespace string) *daemonSet {
	return &daemonSet{cli: cli, ns: namespace}
}
