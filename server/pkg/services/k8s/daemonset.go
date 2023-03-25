package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type DaemonSetImp interface {
	List(ctx context.Context) ([]*v1.DaemonSet, error)
	Get(ctx context.Context, name string) (*v1.DaemonSet, error)
	Create(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error)
	Update(ctx context.Context, obj *v1.DaemonSet) (*v1.DaemonSet, error)
	Delete(ctx context.Context, name string) error
}

type daemonSet struct {
	cli *store.Clients
	ns  string
}

func newDaemonSet(cli *store.Clients, ns string) *daemonSet {
	return &daemonSet{cli: cli, ns: ns}
}

func (d *daemonSet) List(ctx context.Context) ([]*v1.DaemonSet, error) {
	list, err := d.cli.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *daemonSet) Get(ctx context.Context, name string) (*v1.DaemonSet, error) {
	dep, err := d.cli.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.ns).Get(name)
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
