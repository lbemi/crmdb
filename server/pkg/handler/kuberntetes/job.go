package kuberntetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobGetter interface {
	Jobs(namespace string) IJob
}

type IJob interface {
	List(ctx context.Context) (*v1.JobList, error)
	Get(ctx context.Context, name string) (*v1.Job, error)
	Create(ctx context.Context, obj *v1.Job) (*v1.Job, error)
	Update(ctx context.Context, obj *v1.Job) (*v1.Job, error)
	Delete(ctx context.Context, name string) error
}

type job struct {
	cli *cloud.Clients
	ns  string
}

func (d *job) List(ctx context.Context) (*v1.JobList, error) {
	list, err := d.cli.ClientSet.BatchV1().Jobs(d.ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *job) Get(ctx context.Context, name string) (*v1.Job, error) {
	dep, err := d.cli.ClientSet.BatchV1().Jobs(d.ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return dep, err
}

func (d *job) Create(ctx context.Context, obj *v1.Job) (*v1.Job, error) {
	newJob, err := d.cli.ClientSet.BatchV1().Jobs(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return newJob, err
}

func (d *job) Update(ctx context.Context, obj *v1.Job) (*v1.Job, error) {
	updateJob, err := d.cli.ClientSet.BatchV1().Jobs(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return updateJob, err
}

func (d *job) Delete(ctx context.Context, name string) error {
	err := d.cli.ClientSet.BatchV1().Jobs(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func NewJob(cli *cloud.Clients, namespace string) *job {
	return &job{cli: cli, ns: namespace}
}
