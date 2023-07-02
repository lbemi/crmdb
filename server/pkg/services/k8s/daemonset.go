package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"

	"github.com/lbemi/lbemi/pkg/restfulx"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type DaemonSetImp interface {
	List(ctx context.Context) []*v1.DaemonSet
	Get(ctx context.Context, name string) *v1.DaemonSet
	Create(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet
	Update(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet
	Delete(ctx context.Context, name string)
}

type daemonSet struct {
	cli *cache.ClientConfig
	ns  string
}

func newDaemonSet(cli *cache.ClientConfig, ns string) *daemonSet {
	return &daemonSet{cli: cli, ns: ns}
}

func (d *daemonSet) List(ctx context.Context) []*v1.DaemonSet {
	list, err := d.cli.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return list
}

func (d *daemonSet) Get(ctx context.Context, name string) *v1.DaemonSet {
	dep, err := d.cli.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (d *daemonSet) Create(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	newDaemonSet, err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newDaemonSet
}

func (d *daemonSet) Update(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	updateDaemonSet, err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateDaemonSet
}

func (d *daemonSet) Delete(ctx context.Context, name string) {
	err := d.cli.ClientSet.AppsV1().DaemonSets(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}
