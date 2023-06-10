package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
	"strings"
)

type PersistentVolumeClaimGetter interface {
	PersistentVolumeClaim(namespace string) PersistentVolumeClaimImp
}

type PersistentVolumeClaimImp interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.PersistentVolumeClaim
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
	Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
}

type persistentVolumeClaim struct {
	k8s *k8s.Factory
}

func NewPersistentVolumeClaim(k8s *k8s.Factory) *persistentVolumeClaim {
	return &persistentVolumeClaim{k8s: k8s}
}

func (p *persistentVolumeClaim) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := p.k8s.PersistentVolumeClaim().List(ctx)
	res := &form.PageResult{}
	var pvcList = make([]*v1.PersistentVolumeClaim, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				pvcList = append(pvcList, item)
			}
		}
		data = pvcList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				pvcList = append(pvcList, item)
			}
		}
		data = pvcList
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

func (p *persistentVolumeClaim) Get(ctx context.Context, name string) *v1.PersistentVolumeClaim {
	return p.k8s.PersistentVolumeClaim().Get(ctx, name)
}

func (p *persistentVolumeClaim) Delete(ctx context.Context, name string) {
	p.k8s.PersistentVolumeClaim().Delete(ctx, name)
}

func (p *persistentVolumeClaim) Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	return p.k8s.PersistentVolumeClaim().Create(ctx, pvc)
}

func (p *persistentVolumeClaim) Update(ctx context.Context, configMap *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	return p.k8s.PersistentVolumeClaim().Update(ctx, configMap)
}
