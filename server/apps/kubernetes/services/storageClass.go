package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	"k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/labels"
)

type StorageClassGetter interface {
	StorageClass() StorageClassImp
}

type StorageClassImp interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.StorageClass
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, pvc *v1.StorageClass) *v1.StorageClass
}

type StorageClass struct {
	client *cache.ClientConfig
}

func NewStorageClass(client *cache.ClientConfig) *StorageClass {
	return &StorageClass{client: client}
}

func (p *StorageClass) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := p.client.SharedInformerFactory.Storage().V1().StorageClasses().Lister().List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var pvcList = make([]*v1.StorageClass, 0)
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

func (p *StorageClass) Get(ctx context.Context, name string) *v1.StorageClass {
	res, err := p.client.SharedInformerFactory.Storage().V1().StorageClasses().Lister().Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (p *StorageClass) Delete(ctx context.Context, name string) {
	err := p.client.ClientSet.StorageV1().StorageClasses().Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (p *StorageClass) Create(ctx context.Context, pvc *v1.StorageClass) *v1.StorageClass {
	res, err := p.client.ClientSet.StorageV1().StorageClasses().Create(ctx, pvc, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

type StorageClassHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewStorageClassHandler(client *cache.ClientConfig, clusterName string) *StorageClassHandler {
	return &StorageClassHandler{client: client, clusterName: clusterName}
}

func (p *StorageClassHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (p *StorageClassHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (p *StorageClassHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
