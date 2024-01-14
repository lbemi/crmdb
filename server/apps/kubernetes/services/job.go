package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type JobGetter interface {
	Jobs(namespace string) IJob
}

type IJob interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.Job
	GetJobPods(ctx context.Context, name string) []*corev1.Pod
	Create(ctx context.Context, obj *v1.Job) *v1.Job
	Update(ctx context.Context, obj *v1.Job) *v1.Job
	Delete(ctx context.Context, name string)
}

type job struct {
	client    *cache.ClientConfig
	namespace string
}

func NewJob(cli *cache.ClientConfig, namespace string) *job {
	return &job{client: cli, namespace: namespace}
}

func (d *job) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := d.client.SharedInformerFactory.Batch().V1().Jobs().Lister().Jobs(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
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
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				jobList = append(jobList, item)
			}
		}
		data = jobList
	}
	//按时间排序
	sort.SliceStable(data, func(i, j int) bool {
		return data[j].ObjectMeta.GetCreationTimestamp().Time.Before(data[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	total := len(data)
	for _, item := range data {
		util.RestoreGVK(item)
	}
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
	res, err := d.client.SharedInformerFactory.Batch().V1().Jobs().Lister().Jobs(d.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}
func (d *job) GetJobPods(ctx context.Context, name string) []*corev1.Pod {
	dae := d.Get(ctx, name)

	res := make([]*corev1.Pod, 0)
	pods, err := d.client.SharedInformerFactory.Core().V1().Pods().Lister().Pods(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	for _, item := range pods {
		if metav1.IsControlledBy(item, dae) {
			res = append(res, item)
		}
	}

	restoreGVKForList(res)

	sort.Slice(res, func(i, j int) bool {
		// sort by creation timestamp in descending order
		if res[j].ObjectMeta.GetCreationTimestamp().Time.Before(res[i].ObjectMeta.GetCreationTimestamp().Time) {
			return true
		} else if res[i].ObjectMeta.GetCreationTimestamp().Time.Before(res[j].ObjectMeta.GetCreationTimestamp().Time) {
			return false
		}

		// if the creation timestamps are equal, sort by name in ascending order
		return res[i].ObjectMeta.GetName() < res[j].ObjectMeta.GetName()
	})

	return res
}
func (d *job) Create(ctx context.Context, obj *v1.Job) *v1.Job {
	newJob, err := d.client.ClientSet.BatchV1().Jobs(d.namespace).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newJob
}

func (d *job) Update(ctx context.Context, obj *v1.Job) *v1.Job {
	updateJob, err := d.client.ClientSet.BatchV1().Jobs(d.namespace).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateJob
}

func (d *job) Delete(ctx context.Context, name string) {
	err := d.client.ClientSet.BatchV1().Jobs(d.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}
