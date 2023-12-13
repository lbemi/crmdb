package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/labels"
)

type TaskGetter interface {
	Tasks(namespace string) ITask
}

type ITask interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.Task
	Create(ctx context.Context, obj *v1.Task) *v1.Task
	Update(ctx context.Context, obj *v1.Task) *v1.Task
	Delete(ctx context.Context, name string)
}

type task struct {
	cli *cache.ClientConfig
	ns  string
}

func NewTask(client *cache.ClientConfig, namespace string) ITask {
	return &task{cli: client, ns: namespace}
}

func (t *task) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := t.cli.TektonSharedInformerFactory.Tekton().V1().Tasks().Lister().Tasks(t.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var taskMapList = make([]*v1.Task, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				taskMapList = append(taskMapList, item)
			}
		}
		data = taskMapList
	}
	//按时间排序
	sort.SliceStable(taskMapList, func(i, j int) bool {
		return taskMapList[j].ObjectMeta.GetCreationTimestamp().Time.Before(taskMapList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	if label != "" {
		for _, item := range taskMapList {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				taskMapList = append(taskMapList, item)
			}
		}
		data = taskMapList
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

func (t *task) Get(ctx context.Context, name string) *v1.Task {
	res, err := t.cli.TektonSharedInformerFactory.Tekton().V1().Tasks().Lister().Tasks(t.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (t *task) Create(ctx context.Context, obj *v1.Task) *v1.Task {
	newTask, err := t.cli.TektonClient.TektonV1().Tasks(t.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newTask
}

func (t *task) Update(ctx context.Context, obj *v1.Task) *v1.Task {
	updateTask, err := t.cli.TektonClient.TektonV1().Tasks(t.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateTask
}

func (t *task) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(t.cli.TektonClient.TektonV1().Tasks(t.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}
