package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type JobImp interface {
	List(ctx context.Context) ([]*v1.Job, error)
	Get(ctx context.Context, name string) (*v1.Job, error)
	Create(ctx context.Context, obj *v1.Job) (*v1.Job, error)
	Update(ctx context.Context, obj *v1.Job) (*v1.Job, error)
	Delete(ctx context.Context, name string) error
}

type job struct {
	cli *store.Clients
	ns  string
}

func (d *job) List(ctx context.Context) ([]*v1.Job, error) {
	list, err := d.cli.SharedInformerFactory.Batch().V1().Jobs().Lister().Jobs(d.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	return list, err
}

func (d *job) Get(ctx context.Context, name string) (*v1.Job, error) {
	job, err := d.cli.SharedInformerFactory.Batch().V1().Jobs().Lister().Jobs(d.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return job, err
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

func newJob(cli *store.Clients, namespace string) *job {
	return &job{cli: cli, ns: namespace}
}
