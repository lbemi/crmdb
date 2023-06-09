package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type CronJobImp interface {
	List(ctx context.Context) []*v1.CronJob
	Get(ctx context.Context, name string) *v1.CronJob
	Create(ctx context.Context, obj *v1.CronJob) *v1.CronJob
	Update(ctx context.Context, obj *v1.CronJob) *v1.CronJob
	Delete(ctx context.Context, name string)
}

type cronJob struct {
	cli *store.ClientConfig
	ns  string
}

func (d *cronJob) List(ctx context.Context) []*v1.CronJob {
	list, err := d.cli.SharedInformerFactory.Batch().V1().CronJobs().Lister().CronJobs(d.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return list
}

func (d *cronJob) Get(ctx context.Context, name string) *v1.CronJob {
	dep, err := d.cli.SharedInformerFactory.Batch().V1().CronJobs().Lister().CronJobs(d.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (d *cronJob) Create(ctx context.Context, obj *v1.CronJob) *v1.CronJob {
	newCronJob, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newCronJob
}

func (d *cronJob) Update(ctx context.Context, obj *v1.CronJob) *v1.CronJob {
	updateCronJob, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateCronJob
}

func (d *cronJob) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(d.cli.ClientSet.BatchV1().CronJobs(d.ns).
		Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func newCronJob(cli *store.ClientConfig, namespace string) *cronJob {
	return &cronJob{cli: cli, ns: namespace}
}
