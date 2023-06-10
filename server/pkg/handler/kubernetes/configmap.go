package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"
	v1 "k8s.io/api/core/v1"
	"strings"
)

type ConfigMapGetter interface {
	ConfigMaps(namespace string) IConfigMap
}

type IConfigMap interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.ConfigMap
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.ConfigMap) *v1.ConfigMap
	Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap
}

type configMap struct {
	k8s *k8s.Factory
}

func NewConfigMap(k8s *k8s.Factory) *configMap {
	return &configMap{k8s: k8s}
}

func (s *configMap) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := s.k8s.ConfigMap().List(ctx)
	res := &form.PageResult{}

	var configMapList = make([]*v1.ConfigMap, 0)

	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				configMapList = append(configMapList, item)
			}
		}
		data = configMapList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(item.Name, label) {
				configMapList = append(configMapList, item)
			}
		}
		data = configMapList
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

func (s *configMap) Get(ctx context.Context, name string) *v1.ConfigMap {
	return s.k8s.ConfigMap().Get(ctx, name)
}

func (s *configMap) Delete(ctx context.Context, name string) {
	s.k8s.ConfigMap().Delete(ctx, name)
}

func (s *configMap) Create(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	return s.k8s.ConfigMap().Create(ctx, configMap)
}

func (s *configMap) Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	return s.k8s.ConfigMap().Update(ctx, configMap)
}
