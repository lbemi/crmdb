package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/apps/v1"
	"strings"
)

type StatefulSetGetter interface {
	StatefulSets(namespace string) IStatefulSet
}

type IStatefulSet interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.StatefulSet
	Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet
	Delete(ctx context.Context, name string)
}

type statefulSet struct {
	k8s *k8s.Factory
}

func NewStatefulSet(k8s *k8s.Factory) *statefulSet {
	return &statefulSet{k8s: k8s}
}

func (d *statefulSet) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := d.k8s.StatefulSet().List(ctx)
	res := &form.PageResult{}
	var statefulSetList = make([]*v1.StatefulSet, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				statefulSetList = append(statefulSetList, item)
			}
		}
		data = statefulSetList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				statefulSetList = append(statefulSetList, item)
			}
		}
		data = statefulSetList
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

func (d *statefulSet) Get(ctx context.Context, name string) *v1.StatefulSet {
	return d.k8s.StatefulSet().Get(ctx, name)
}

func (d *statefulSet) Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	return d.k8s.StatefulSet().Create(ctx, obj)
}

func (d *statefulSet) Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	return d.k8s.StatefulSet().Update(ctx, obj)
}

func (d *statefulSet) Delete(ctx context.Context, name string) {
	d.k8s.StatefulSet().Delete(ctx, name)
}
