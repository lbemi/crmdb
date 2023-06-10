package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/batch/v1"
	"strings"
)

type JobGetter interface {
	Jobs(namespace string) IJob
}

type IJob interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
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

func (d *job) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := d.k8s.Job().List(ctx)
	res := &form.PageResult{}
	var jobList = make([]*v1.Job, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				jobList = append(jobList, item)
			}
		}
		data = jobList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				jobList = append(jobList, item)
			}
		}
		data = jobList
	}

	total := len(data)
	// 未传递分页查询参数
	if query.Limit == 0 && query.Page == 0 {
		res.Data = data
	} else {
		if total <= query.Limit {
			res.Data = data
		} else if query.Page*query.Limit >= total {
			res.Data = data[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = data[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}
	res.Total = int64(total)
	return res
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
