package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type StatefulSetImp interface {
	List(ctx context.Context) []*v1.StatefulSet
	Get(ctx context.Context, name string) *v1.StatefulSet
	Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Delete(ctx context.Context, name string)
}

type statefulSet struct {
	cli *store.ClientConfig
	ns  string
}

func (d *statefulSet) List(ctx context.Context) []*v1.StatefulSet {
	list, err := d.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(d.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return list
}

func (d *statefulSet) Get(ctx context.Context, name string) *v1.StatefulSet {
	sts, err := d.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(d.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return sts
}

func (d *statefulSet) Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	newStatefulSet, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newStatefulSet
}

func (d *statefulSet) Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	updateStatefulSet, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateStatefulSet
}

func (d *statefulSet) Delete(ctx context.Context, name string) {
	err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func newStatefulSet(cli *store.ClientConfig, namespace string) *statefulSet {
	return &statefulSet{cli: cli, ns: namespace}
}

type StatefulSetHandle struct {
}

func NewStatefulSetHandle() *StatefulSetHandle {
	return &StatefulSetHandle{}
}

func (s *StatefulSetHandle) OnAdd(obj interface{}) {
	//TODO implement me
}

func (s *StatefulSetHandle) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (s *StatefulSetHandle) OnDelete(obj interface{}) {
	//TODO implement me
}
