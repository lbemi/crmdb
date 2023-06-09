package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/batch/v1"
)

type CronJobGetter interface {
	CronJobs(namespace string) ICronJob
}

type ICronJob interface {
	List(ctx context.Context) []*v1.CronJob
	Get(ctx context.Context, name string) *v1.CronJob
	Create(ctx context.Context, obj *v1.CronJob) *v1.CronJob
	Update(ctx context.Context, obj *v1.CronJob) *v1.CronJob
	Delete(ctx context.Context, name string)
}

type cronJob struct {
	k8s *k8s.Factory
}

func NewCronJob(k8s *k8s.Factory) *cronJob {
	return &cronJob{k8s: k8s}
}

func (d *cronJob) List(ctx context.Context) []*v1.CronJob {
	return d.k8s.CronJob().List(ctx)
}

func (d *cronJob) Get(ctx context.Context, name string) *v1.CronJob {
	return d.k8s.CronJob().Get(ctx, name)
}

func (d *cronJob) Create(ctx context.Context, obj *v1.CronJob) *v1.CronJob {
	return d.k8s.CronJob().Create(ctx, obj)
}

func (d *cronJob) Update(ctx context.Context, obj *v1.CronJob) *v1.CronJob {
	return d.k8s.CronJob().Update(ctx, obj)
}

func (d *cronJob) Delete(ctx context.Context, name string) {
	d.k8s.CronJob().Delete(ctx, name)
}
