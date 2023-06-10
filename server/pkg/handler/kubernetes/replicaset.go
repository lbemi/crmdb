package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	appsv1 "k8s.io/api/apps/v1"
	"strings"
)

type ReplicasetGetter interface {
	Replicaset(namespace string) ReplicasetImp
}

type ReplicasetImp interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *appsv1.ReplicaSet
}

type replicaset struct {
	k8s *k8s.Factory
}

func NewReplicaset(k8s *k8s.Factory) *replicaset {
	return &replicaset{k8s: k8s}
}

func (r *replicaset) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := r.k8s.Replicaset().List(ctx)
	res := &form.PageResult{}
	var replicasetList = make([]*appsv1.ReplicaSet, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				replicasetList = append(replicasetList, item)
			}
		}
		data = replicasetList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				replicasetList = append(replicasetList, item)
			}
		}
		data = replicasetList
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

func (r *replicaset) Get(ctx context.Context, name string) *appsv1.ReplicaSet {
	return r.k8s.Replicaset().Get(ctx, name)
}
