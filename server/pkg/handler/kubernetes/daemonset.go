package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/apps/v1"
	"strings"
)

type DaemonSetGetter interface {
	DaemonSets(namespace string) IDaemonSet
}

type IDaemonSet interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.DaemonSet
	Create(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet
	Update(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet
	Delete(ctx context.Context, name string)
}

type daemonSet struct {
	k8s *k8s.Factory
}

func NewDaemonSet(k8s *k8s.Factory) *daemonSet {
	return &daemonSet{k8s: k8s}
}

func (d *daemonSet) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := d.k8s.DaemonSet().List(ctx)
	res := &form.PageResult{}
	var daemonSetList = make([]*v1.DaemonSet, 0)

	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				daemonSetList = append(daemonSetList, item)
			}
		}
		data = daemonSetList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				daemonSetList = append(daemonSetList, item)
			}
		}
		data = daemonSetList
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

func (d *daemonSet) Get(ctx context.Context, name string) *v1.DaemonSet {
	return d.k8s.DaemonSet().Get(ctx, name)
}

func (d *daemonSet) Create(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	return d.k8s.DaemonSet().Create(ctx, obj)
}

func (d *daemonSet) Update(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	return d.k8s.DaemonSet().Update(ctx, obj)
}

func (d *daemonSet) Delete(ctx context.Context, name string) {
	d.k8s.DaemonSet().Delete(ctx, name)
}
