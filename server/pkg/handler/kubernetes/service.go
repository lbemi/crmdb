package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
	"strings"
)

type ServiceGetter interface {
	Service(namespace string) IService
}

type IService interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.Service
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Service) *v1.Service
	Update(ctx context.Context, service *v1.Service) *v1.Service
	ListWorkLoad(ctx context.Context, name string) *types.ServiceWorkLoad
}

type service struct {
	k8s *k8s.Factory
}

func NewService(k8s *k8s.Factory) *service {
	return &service{k8s: k8s}
}

func (s *service) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := s.k8s.Service().List(ctx)
	res := &form.PageResult{}
	var serviceList = make([]*v1.Service, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				serviceList = append(serviceList, item)
			}
		}
		data = serviceList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				serviceList = append(serviceList, item)
			}
		}
		data = serviceList
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

func (s *service) ListWorkLoad(ctx context.Context, name string) *types.ServiceWorkLoad {
	return s.k8s.Service().ListWorkLoad(ctx, name)
}

func (s *service) Get(ctx context.Context, name string) *v1.Service {
	return s.k8s.Service().Get(ctx, name)
}

func (s *service) Delete(ctx context.Context, name string) {
	s.k8s.Service().Delete(ctx, name)
}

func (s *service) Create(ctx context.Context, service *v1.Service) *v1.Service {
	return s.k8s.Service().Create(ctx, service)
}

func (s *service) Update(ctx context.Context, service *v1.Service) *v1.Service {
	return s.k8s.Service().Update(ctx, service)
}
