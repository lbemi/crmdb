package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type PersistentVolumeGetter interface {
	PersistentVolume() PersistentVolumeImp
}

type PersistentVolumeImp interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.PersistentVolume
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, pvc *v1.PersistentVolume) *v1.PersistentVolume
}

type PersistentVolume struct {
	client *cache.ClientConfig
}

func NewPersistentVolume(client *cache.ClientConfig) *PersistentVolume {
	return &PersistentVolume{client: client}
}

func (p *PersistentVolume) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := p.client.SharedInformerFactory.Core().V1().PersistentVolumes().Lister().List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var pvcList = make([]*v1.PersistentVolume, 0)
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
	for _, item := range data {
		util.RestoreGVK(item)
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

func (p *PersistentVolume) Get(ctx context.Context, name string) *v1.PersistentVolume {
	res, err := p.client.SharedInformerFactory.Core().V1().PersistentVolumes().Lister().Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (p *PersistentVolume) Delete(ctx context.Context, name string) {
	err := p.client.ClientSet.CoreV1().PersistentVolumes().Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (p *PersistentVolume) Create(ctx context.Context, pvc *v1.PersistentVolume) *v1.PersistentVolume {
	res, err := p.client.ClientSet.CoreV1().PersistentVolumes().Create(ctx, pvc, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

type PersistentVolumeHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewPersistentVolumeHandler(client *cache.ClientConfig, clusterName string) *PersistentVolumeHandler {
	return &PersistentVolumeHandler{client: client, clusterName: clusterName}
}

func (p *PersistentVolumeHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (p *PersistentVolumeHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (p *PersistentVolumeHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
