package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/labels"
)

type TaskRunGetter interface {
	TaskRuns(namespace string) ITaskRun
}

type ITaskRun interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.TaskRun
	Create(ctx context.Context, obj *v1.TaskRun) *v1.TaskRun
	Update(ctx context.Context, obj *v1.TaskRun) *v1.TaskRun
	Delete(ctx context.Context, name string)
}

type TaskRun struct {
	cli *cache.ClientConfig
	ns  string
}

func NewTaskRun(client *cache.ClientConfig, namespace string) ITaskRun {
	return &TaskRun{cli: client, ns: namespace}
}

func (p *TaskRun) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := p.cli.TektonSharedInformerFactory.Tekton().V1().TaskRuns().Lister().TaskRuns(p.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var TaskRunMapList = make([]*v1.TaskRun, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				TaskRunMapList = append(TaskRunMapList, item)
			}
		}
		data = TaskRunMapList
	}
	//按时间排序
	sort.SliceStable(TaskRunMapList, func(i, j int) bool {
		return TaskRunMapList[j].ObjectMeta.GetCreationTimestamp().Time.Before(TaskRunMapList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	if label != "" {
		for _, item := range TaskRunMapList {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				TaskRunMapList = append(TaskRunMapList, item)
			}
		}
		data = TaskRunMapList
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

func (p *TaskRun) Get(ctx context.Context, name string) *v1.TaskRun {
	res, err := p.cli.TektonSharedInformerFactory.Tekton().V1().TaskRuns().Lister().TaskRuns(p.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (p *TaskRun) Create(ctx context.Context, obj *v1.TaskRun) *v1.TaskRun {
	newTaskRun, err := p.cli.TektonClient.TektonV1().TaskRuns(p.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newTaskRun
}

func (p *TaskRun) Update(ctx context.Context, obj *v1.TaskRun) *v1.TaskRun {
	updateTaskRun, err := p.cli.TektonClient.TektonV1().TaskRuns(p.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateTaskRun
}

func (p *TaskRun) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(p.cli.TektonClient.TektonV1().TaskRuns(p.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}
