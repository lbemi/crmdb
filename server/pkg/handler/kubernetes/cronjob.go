package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/batch/v1"
	"strings"
)

type CronJobGetter interface {
	CronJobs(namespace string) ICronJob
}

type ICronJob interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
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

func (d *cronJob) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := d.k8s.CronJob().List(ctx)
	res := &form.PageResult{}
	var cronJobMapList = make([]*v1.CronJob, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				cronJobMapList = append(cronJobMapList, item)
			}
		}
		data = cronJobMapList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				cronJobMapList = append(cronJobMapList, item)
			}
		}
		data = cronJobMapList
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
