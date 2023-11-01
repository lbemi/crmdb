package kubernetes

import (
	"context"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
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
	client    *store.ClientConfig
	namespace string
}

func NewDaemonSet(client *store.ClientConfig, namespace string) *daemonSet {
	return &daemonSet{client: client, namespace: namespace}
}

func (d *daemonSet) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data, err := d.client.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
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
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				daemonSetList = append(daemonSetList, item)
			}
		}
		data = daemonSetList
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

func (d *daemonSet) Get(ctx context.Context, name string) *v1.DaemonSet {
	dep, err := d.client.SharedInformerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(d.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (d *daemonSet) Create(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	newDaemonSet, err := d.client.ClientSet.AppsV1().DaemonSets(d.namespace).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newDaemonSet
}

func (d *daemonSet) Update(ctx context.Context, obj *v1.DaemonSet) *v1.DaemonSet {
	updateDaemonSet, err := d.client.ClientSet.AppsV1().DaemonSets(d.namespace).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateDaemonSet
}

func (d *daemonSet) Delete(ctx context.Context, name string) {
	err := d.client.ClientSet.AppsV1().DaemonSets(d.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}
