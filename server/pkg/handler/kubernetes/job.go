package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/batch/v1"
)

type JobGetter interface {
	Jobs(namespace string) IJob
}

type IJob interface {
	List(ctx context.Context) []*v1.Job
	Get(ctx context.Context, name string) *v1.Job
	Create(ctx context.Context, obj *v1.Job) *v1.Job
	Update(ctx context.Context, obj *v1.Job) *v1.Job
	Delete(ctx context.Context, name string)
}

type job struct {
	k8s *k8s.Factory
}

func NewJob(k8s *k8s.Factory) *job {
	return &job{k8s: k8s}
}

func (d *job) List(ctx context.Context) []*v1.Job {
	return d.k8s.Job().List(ctx)
}

func (d *job) Get(ctx context.Context, name string) *v1.Job {
	return d.k8s.Job().Get(ctx, name)
}

func (d *job) Create(ctx context.Context, obj *v1.Job) *v1.Job {
	return d.k8s.Job().Create(ctx, obj)
}

func (d *job) Update(ctx context.Context, obj *v1.Job) *v1.Job {
	return d.k8s.Job().Update(ctx, obj)
}

func (d *job) Delete(ctx context.Context, name string) {
	d.k8s.Job().Delete(ctx, name)
}
