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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
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

type PersistentVolumeClaim struct {
	client *store.ClientConfig
	ns     string
}

func NewPersistentVolumeClaim(client *store.ClientConfig, namespace string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{client: client, ns: namespace}
}

func (p *PersistentVolumeClaim) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data, err := p.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(p.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
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
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				pvcList = append(pvcList, item)
			}
		}
		data = pvcList
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

func (p *PersistentVolumeClaim) Get(ctx context.Context, name string) *v1.PersistentVolumeClaim {
	res, err := p.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(p.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (p *PersistentVolumeClaim) Delete(ctx context.Context, name string) {
	err := p.client.ClientSet.CoreV1().PersistentVolumeClaims(p.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (p *PersistentVolumeClaim) Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	res, err := p.client.ClientSet.CoreV1().PersistentVolumeClaims(p.ns).Create(ctx, pvc, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (p *PersistentVolumeClaim) Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	res, err := p.client.ClientSet.CoreV1().PersistentVolumeClaims(p.ns).Update(ctx, pvc, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

type PersistentVolumeClaimHandler struct {
	client      *store.ClientConfig
	clusterName string
}

func NewPersistentVolumeClaimHandler(client *store.ClientConfig, clusterName string) *PersistentVolumeClaimHandler {
	return &PersistentVolumeClaimHandler{client: client, clusterName: clusterName}
}

func (p *PersistentVolumeClaimHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (p *PersistentVolumeClaimHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (p *PersistentVolumeClaimHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
