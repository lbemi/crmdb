package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CronJobGetter interface {
	CronJobs(namespace string) ICronJob
}

type ICronJob interface {
	List(ctx context.Context) (*v1.CronJobList, error)
	Get(ctx context.Context, name string) (*v1.CronJob, error)
	Create(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error)
	Update(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error)
	Delete(ctx context.Context, name string) error
}

type cronJob struct {
	cli *cloud.Clients
	ns  string
}

func (d *cronJob) List(ctx context.Context) (*v1.CronJobList, error) {
	list, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *cronJob) Get(ctx context.Context, name string) (*v1.CronJob, error) {
	dep, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Get(ctx, name, metav1.GetOptions{})
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

func NewCronJob(cli *cloud.Clients, namespace string) *cronJob {
	return &cronJob{cli: cli, ns: namespace}
}
