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

type StatefulSet struct {
	cli *store.ClientConfig
	ns  string
}

func NewStatefulSet(client *store.ClientConfig, ns string) *StatefulSet {
	return &StatefulSet{cli: client, ns: ns}
}

func (d *StatefulSet) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data, err := d.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(d.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
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
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				statefulSetList = append(statefulSetList, item)
			}
		}
		data = statefulSetList
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

func (d *StatefulSet) Get(ctx context.Context, name string) *v1.StatefulSet {
	sts, err := d.cli.SharedInformerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(d.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return sts
}

func (d *StatefulSet) Create(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	newStatefulSet, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newStatefulSet
}

func (d *StatefulSet) Update(ctx context.Context, obj *v1.StatefulSet) *v1.StatefulSet {
	updateStatefulSet, err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateStatefulSet
}

func (d *StatefulSet) Delete(ctx context.Context, name string) {
	err := d.cli.ClientSet.AppsV1().StatefulSets(d.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

type StatefulSetHandler struct {
}

func NewStatefulSetHandle() *StatefulSetHandler {
	return &StatefulSetHandler{}
}

func (s *StatefulSetHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (s *StatefulSetHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (s *StatefulSetHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
