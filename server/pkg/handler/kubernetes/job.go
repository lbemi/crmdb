package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/batch/v1"
)

type JobGetter interface {
	Jobs(namespace string) IJob
}

type IJob interface {
	List(ctx context.Context) ([]*v1.Job, error)
	Get(ctx context.Context, name string) (*v1.Job, error)
	Create(ctx context.Context, obj *v1.Job) (*v1.Job, error)
	Update(ctx context.Context, obj *v1.Job) (*v1.Job, error)
	Delete(ctx context.Context, name string) error
}

type job struct {
	k8s *k8s.Factory
}

func NewJob(k8s *k8s.Factory) *job {
	return &job{k8s: k8s}
}

func (d *job) List(ctx context.Context) ([]*v1.Job, error) {
	list, err := d.k8s.Job().List(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *job) Get(ctx context.Context, name string) (*v1.Job, error) {
	job, err := d.k8s.Job().Get(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return job, err
}

func (d *job) Create(ctx context.Context, obj *v1.Job) (*v1.Job, error) {
	newJob, err := d.k8s.Job().Create(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return newJob, err
}

func (d *job) Update(ctx context.Context, obj *v1.Job) (*v1.Job, error) {
	updateJob, err := d.k8s.Job().Update(ctx, obj)
	if err != nil {
		log.Logger.Error(err)
	}
	return updateJob, err
}

func (d *job) Delete(ctx context.Context, name string) error {
	err := d.k8s.Job().Delete(ctx, name)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}
