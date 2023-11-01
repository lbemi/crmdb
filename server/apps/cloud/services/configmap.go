package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ConfigMapGetter interface {
	ConfigMaps(namespace string) IConfigMap
}

type IConfigMap interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageConfigMap
	Get(ctx context.Context, name string) *v1.ConfigMap
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.ConfigMap) *v1.ConfigMap
	Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap
}

type ConfigMap struct {
	client    *cache.ClientConfig
	namespace string
}

func NewConfigMap(client *cache.ClientConfig, namespace string) *ConfigMap {
	return &ConfigMap{client: client, namespace: namespace}
}

func (c *ConfigMap) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageConfigMap {
	data, err := c.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(c.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageConfigMap{}
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
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				configMapList = append(configMapList, item)
			}
		}
		data = configMapList
	}

	//按时间排序
	sort.SliceStable(data, func(i, j int) bool {
		return data[j].ObjectMeta.GetCreationTimestamp().Time.Before(data[i].ObjectMeta.GetCreationTimestamp().Time)
	})
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

func (c *ConfigMap) Get(ctx context.Context, name string) *v1.ConfigMap {
	res, err := c.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(c.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (c *ConfigMap) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(
		c.client.ClientSet.CoreV1().ConfigMaps(c.namespace).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func (c *ConfigMap) Create(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	log.Logger.Error(c.namespace)
	res, err := c.client.ClientSet.CoreV1().ConfigMaps(c.namespace).Create(ctx, configMap, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (c *ConfigMap) Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	res, err := c.client.ClientSet.CoreV1().ConfigMaps(c.namespace).Update(ctx, configMap, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}
