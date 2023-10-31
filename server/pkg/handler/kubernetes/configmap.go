package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ConfigMapGetter interface {
	ConfigMaps(namespace string) IConfigMap
}

type IConfigMap interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageConfigMap
	Get(ctx context.Context, name string) *v1.ConfigMap
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.ConfigMap) *v1.ConfigMap
	Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap
}

type configMap struct {
	client *store.ClientConfig
	ns     string
}

func NewConfigMap(client *store.ClientConfig, namespace string) *configMap {
	return &configMap{client: client, ns: namespace}
}

func (c *configMap) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageConfigMap {
	data, err := c.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(c.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &form.PageConfigMap{}
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

func (c *configMap) Get(ctx context.Context, name string) *v1.ConfigMap {
	res, err := c.client.SharedInformerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(c.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (c *configMap) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(
		c.client.ClientSet.CoreV1().ConfigMaps(c.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func (c *configMap) Create(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	log.Logger.Error(c.ns)
	res, err := c.client.ClientSet.CoreV1().ConfigMaps(c.ns).Create(ctx, configMap, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (c *configMap) Update(ctx context.Context, configMap *v1.ConfigMap) *v1.ConfigMap {
	res, err := c.client.ClientSet.CoreV1().ConfigMaps(c.ns).Update(ctx, configMap, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}
