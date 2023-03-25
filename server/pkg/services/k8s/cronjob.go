package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type CronJobImp interface {
	List(ctx context.Context) ([]*v1.CronJob, error)
	Get(ctx context.Context, name string) (*v1.CronJob, error)
	Create(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error)
	Update(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error)
	Delete(ctx context.Context, name string) error
}

type cronJob struct {
	cli *store.Clients
	ns  string
}

func (d *cronJob) List(ctx context.Context) ([]*v1.CronJob, error) {
	list, err := d.cli.SharedInformerFactory.Batch().V1().CronJobs().Lister().CronJobs(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *cronJob) Get(ctx context.Context, name string) (*v1.CronJob, error) {
	dep, err := d.cli.SharedInformerFactory.Batch().V1().CronJobs().Lister().CronJobs(d.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *cronJob) Create(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error) {
	newCronJob, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newCronJob, err
}

func (d *cronJob) Update(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error) {
	updateCronJob, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updateCronJob, err
}

func (d *cronJob) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func newCronJob(cli *store.Clients, namespace string) *cronJob {
	return &cronJob{cli: cli, ns: namespace}
}
