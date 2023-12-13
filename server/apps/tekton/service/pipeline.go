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

type PipelineGetter interface {
	Pipelines(namespace string) IPipeline
}

type IPipeline interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.Pipeline
	Create(ctx context.Context, obj *v1.Pipeline) *v1.Pipeline
	Update(ctx context.Context, obj *v1.Pipeline) *v1.Pipeline
	Delete(ctx context.Context, name string)
}

type Pipeline struct {
	cli *cache.ClientConfig
	ns  string
}

func NewPipeline(client *cache.ClientConfig, namespace string) IPipeline {
	return &Pipeline{cli: client, ns: namespace}
}

func (p *Pipeline) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := p.cli.TektonSharedInformerFactory.Tekton().V1().Pipelines().Lister().Pipelines(p.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var PipelineMapList = make([]*v1.Pipeline, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				PipelineMapList = append(PipelineMapList, item)
			}
		}
		data = PipelineMapList
	}
	//按时间排序
	sort.SliceStable(PipelineMapList, func(i, j int) bool {
		return PipelineMapList[j].ObjectMeta.GetCreationTimestamp().Time.Before(PipelineMapList[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	if label != "" {
		for _, item := range PipelineMapList {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				PipelineMapList = append(PipelineMapList, item)
			}
		}
		data = PipelineMapList
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

func (p *Pipeline) Get(ctx context.Context, name string) *v1.Pipeline {
	res, err := p.cli.TektonSharedInformerFactory.Tekton().V1().Pipelines().Lister().Pipelines(p.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (p *Pipeline) Create(ctx context.Context, obj *v1.Pipeline) *v1.Pipeline {
	newPipeline, err := p.cli.TektonClient.TektonV1().Pipelines(p.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newPipeline
}

func (p *Pipeline) Update(ctx context.Context, obj *v1.Pipeline) *v1.Pipeline {
	updatePipeline, err := p.cli.TektonClient.TektonV1().Pipelines(p.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updatePipeline
}

func (p *Pipeline) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(p.cli.TektonClient.TektonV1().Pipelines(p.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}
