package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSetGetter interface {
	StatefulSets(namespace string) IStatefulSet
}

type IStatefulSet interface {
	List(ctx context.Context) (*v1.StatefulSetList, error)
	Get(ctx context.Context, name string) (*v1.StatefulSet, error)
	Create(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error)
	Update(ctx context.Context, obj *v1.StatefulSet) (*v1.StatefulSet, error)
	Delete(ctx context.Context, name string) error
}

type statefulSet struct {
	cli *cloud.Clients
	ns  string
}

func (d *statefulSet) List(ctx context.Context) (*v1.StatefulSetList, error) {
	list, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).List(ctx, metav1.ListOptions{
		Limit: 2,
	})
	if err != nil {
		log.Logger.Error(err)
	}

	return list, err
}

func (d *statefulSet) Get(ctx context.Context, name string) (*v1.StatefulSet, error) {
	dep, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
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

func NewStatefulSet(cli *cloud.Clients, namespace string) *statefulSet {
	return &statefulSet{cli: cli, ns: namespace}
}
