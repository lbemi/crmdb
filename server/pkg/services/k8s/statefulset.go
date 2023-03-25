package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type StatefulSetImp interface {
	List(ctx context.Context) ([]*v1.StatefulSet, error)
	Get(ctx context.Context, name string) (*v1.StatefulSet, error)
	Create(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error)
	Update(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error)
	Delete(ctx context.Context, name string) error
}

type statefulSet struct {
	cli *store.Clients
	ns  string
}

func (d *statefulSet) List(ctx context.Context) ([]*v1.StatefulSet, error) {
	list, err := d.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}

	return list, err
}

func (d *statefulSet) Get(ctx context.Context, name string) (*v1.StatefulSet, error) {
	sts, err := d.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(d.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return sts, err
}

func (d *statefulSet) Create(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error) {
	newStatefulSet, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newStatefulSet, err
}

func (d *statefulSet) Update(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error) {
	updateStatefulSet, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updateStatefulSet, err
}

func (d *statefulSet) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func newStatefulSet(cli *store.Clients, namespace string) *statefulSet {
	return &statefulSet{cli: cli, ns: namespace}
}
