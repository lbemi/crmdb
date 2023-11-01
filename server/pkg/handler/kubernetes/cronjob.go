package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/labels"
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
	cli *store.ClientConfig
	ns  string
}

func NewCronJob(client *store.ClientConfig, namespace string) *cronJob {
	return &cronJob{cli: client, ns: namespace}
}

func (d *cronJob) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data, err := d.cli.SharedInformerFactory.Batch().V1().CronJobs().Lister().CronJobs(d.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
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
	//按时间排序
	sort.SliceStable(data, func(i, j int) bool {
		return data[j].ObjectMeta.GetCreationTimestamp().Time.Before(data[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
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
	dep, err := d.cli.SharedInformerFactory.Batch().V1().CronJobs().Lister().CronJobs(d.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (d *cronJob) Create(ctx context.Context, obj *v1.CronJob) *v1.CronJob {
	newCronJob, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newCronJob
}

func (d *cronJob) Update(ctx context.Context, obj *v1.CronJob) *v1.CronJob {
	updateCronJob, err := d.cli.ClientSet.BatchV1().CronJobs(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateCronJob
}

func (d *cronJob) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(d.cli.ClientSet.BatchV1().CronJobs(d.ns).
		Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}
