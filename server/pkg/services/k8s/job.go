package k8s

import (
	"context"
	"sort"

	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"

	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type JobImp interface {
	List(ctx context.Context) []*v1.Job
	Get(ctx context.Context, name string) *v1.Job
	Create(ctx context.Context, obj *v1.Job) *v1.Job
	Update(ctx context.Context, obj *v1.Job) *v1.Job
	Delete(ctx context.Context, name string)
}

type job struct {
	cli *store.ClientConfig
	ns  string
}

func (d *job) List(ctx context.Context) []*v1.Job {
	list, err := d.cli.SharedInformerFactory.Batch().V1().Jobs().Lister().Jobs(d.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	sort.Slice(list, func(i, j int) bool {
		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
	})
	return list
}

func (d *job) Get(ctx context.Context, name string) *v1.Job {
	res, err := d.cli.SharedInformerFactory.Batch().V1().Jobs().Lister().Jobs(d.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (d *job) Create(ctx context.Context, obj *v1.Job) *v1.Job {
	newJob, err := d.cli.ClientSet.BatchV1().Jobs(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newJob
}

func (d *job) Update(ctx context.Context, obj *v1.Job) *v1.Job {
	updateJob, err := d.cli.ClientSet.BatchV1().Jobs(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateJob
}

func (d *job) Delete(ctx context.Context, name string) {
	err := d.cli.ClientSet.BatchV1().Jobs(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func newJob(cli *store.ClientConfig, namespace string) *job {
	return &job{cli: cli, ns: namespace}
}
