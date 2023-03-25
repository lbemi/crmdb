package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/batch/v1"
)

type CronJobGetter interface {
	CronJobs(namespace string) ICronJob
}

type ICronJob interface {
	List(ctx context.Context) ([]*v1.CronJob, error)
	Get(ctx context.Context, name string) (*v1.CronJob, error)
	Create(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error)
	Update(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error)
	Delete(ctx context.Context, name string) error
}

type cronJob struct {
	k8s *k8s.Factory
}

func NewCronJob(k8s *k8s.Factory) *cronJob {
	return &cronJob{k8s: k8s}
}

func (d *cronJob) List(ctx context.Context) ([]*v1.CronJob, error) {
	list, err := d.k8s.CronJob().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *cronJob) Get(ctx context.Context, name string) (*v1.CronJob, error) {
	dep, err := d.k8s.CronJob().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *cronJob) Create(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error) {
	newCronJob, err := d.k8s.CronJob().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newCronJob, err
}

func (d *cronJob) Update(ctx context.Context, obj *v1.CronJob) (*v1.CronJob, error) {
	updateCronJob, err := d.k8s.CronJob().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return updateCronJob, err
}

func (d *cronJob) Delete(ctx context.Context, name string) error {
	err := d.k8s.CronJob().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}
