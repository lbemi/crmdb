package services

import (
	"context"
	"sort"
	"strings"

	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type VirtualServiceGetter interface {
	VirtualServices(namespace string) IVirtualService
}

type IVirtualService interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageVirtualService
	Get(ctx context.Context, name string) *v1beta1.VirtualService
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1beta1.VirtualService) *v1beta1.VirtualService
	Update(ctx context.Context, VirtualService *v1beta1.VirtualService) *v1beta1.VirtualService
}

type VirtualServices struct {
	cli *cache.ClientConfig
	ns  string
}

func NewVirtualService(cli *cache.ClientConfig, ns string) IVirtualService {
	return &VirtualServices{cli: cli, ns: ns}
}

func (v *VirtualServices) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageVirtualService {
	data, err := v.cli.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Lister().VirtualServices(v.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageVirtualService{}
	var VirtualServiceList = make([]*v1beta1.VirtualService, 0)

	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				VirtualServiceList = append(VirtualServiceList, item)
			}
		}
		data = VirtualServiceList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				VirtualServiceList = append(VirtualServiceList, item)
			}
		}
		data = VirtualServiceList
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

func (v *VirtualServices) Get(ctx context.Context, name string) *v1beta1.VirtualService {
	vs, err := v.cli.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Lister().VirtualServices(v.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return vs
}

func (v *VirtualServices) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(v.cli.IstioClient.NetworkingV1beta1().VirtualServices(v.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func (v *VirtualServices) Create(ctx context.Context, virtualService *v1beta1.VirtualService) *v1beta1.VirtualService {
	newVirtualService, err := v.cli.IstioClient.NetworkingV1beta1().VirtualServices(v.ns).Create(ctx, virtualService, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newVirtualService
}

func (v *VirtualServices) Update(ctx context.Context, virtualService *v1beta1.VirtualService) *v1beta1.VirtualService {
	updateVirtualService, err := v.cli.IstioClient.NetworkingV1beta1().VirtualServices(v.ns).Update(ctx, virtualService, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateVirtualService
}

type VirtualServiceHandler struct {
	clusterName string
	client      *cache.ClientConfig
}

func NewVirtualServiceHandler(client *cache.ClientConfig, clusterName string) *VirtualServiceHandler {
	return &VirtualServiceHandler{client: client, clusterName: clusterName}
}

func (v *VirtualServiceHandler) OnAdd(obj interface{}, isInInitialList bool) {
	v.notifyVirtualServices(obj)
}

func (v *VirtualServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	v.notifyVirtualServices(newObj)
}

func (v *VirtualServiceHandler) OnDelete(obj interface{}) {
	v.notifyVirtualServices(obj)
}

func (v *VirtualServiceHandler) notifyVirtualServices(obj interface{}) {
	namespace := obj.(*v1beta1.VirtualService).Namespace
	VirtualServices, err := v.client.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Lister().VirtualServices(namespace).List(labels.Everything())
	if err != nil {
		global.Logger.Error(err)
	}

	//按时间排序
	sort.SliceStable(VirtualServices, func(i, j int) bool {
		return VirtualServices[j].ObjectMeta.GetCreationTimestamp().Time.Before(VirtualServices[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	go cache.WebsocketStore.SendClusterResource(v.clusterName, "virtualService", map[string]interface{}{
		"cluster": v.clusterName,
		"type":    "virtualService",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      VirtualServices,
		},
	})
}
