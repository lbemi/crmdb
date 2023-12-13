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

type PipelineRunGetter interface {
	PipelineRuns(namespace string) IPipelineRun
}

type IPipelineRun interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.PipelineRun
	Create(ctx context.Context, obj *v1.PipelineRun) *v1.PipelineRun
	Update(ctx context.Context, obj *v1.PipelineRun) *v1.PipelineRun
	Delete(ctx context.Context, name string)
}

type PipelineRun struct {
	cli *cache.ClientConfig
	ns  string
}

func NewPipelineRun(client *cache.ClientConfig, namespace string) IPipelineRun {
	return &PipelineRun{cli: client, ns: namespace}
}

func (p *PipelineRun) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := p.cli.TektonSharedInformerFactory.Tekton().V1().PipelineRuns().Lister().PipelineRuns(p.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var PipelineRunMapList = make([]*v1.PipelineRun, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				PipelineRunMapList = append(PipelineRunMapList, item)
			}
		}
		data = PipelineRunMapList
	}
	//按时间排序
	sort.SliceStable(PipelineRunMapList, func(i, j int) bool {
		return PipelineRunMapList[j].ObjectMeta.GetCreationTimestamp().Time.Before(PipelineRunMapList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	if label != "" {
		for _, item := range PipelineRunMapList {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				PipelineRunMapList = append(PipelineRunMapList, item)
			}
		}
		data = PipelineRunMapList
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

func (p *PipelineRun) Get(ctx context.Context, name string) *v1.PipelineRun {
	res, err := p.cli.TektonSharedInformerFactory.Tekton().V1().PipelineRuns().Lister().PipelineRuns(p.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (p *PipelineRun) Create(ctx context.Context, obj *v1.PipelineRun) *v1.PipelineRun {
	newPipelineRun, err := p.cli.TektonClient.TektonV1().PipelineRuns(p.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newPipelineRun
}

func (p *PipelineRun) Update(ctx context.Context, obj *v1.PipelineRun) *v1.PipelineRun {
	updatePipelineRun, err := p.cli.TektonClient.TektonV1().PipelineRuns(p.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updatePipelineRun
}

func (p *PipelineRun) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(p.cli.TektonClient.TektonV1().PipelineRuns(p.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}
