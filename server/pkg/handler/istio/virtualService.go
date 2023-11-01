package istio

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/restfulx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
)

type VirtualServiceGetter interface {
	VirtualServices(namespace string) IVirtualService
}

type IVirtualService interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageVirtualService
	Get(ctx context.Context, name string) *v1beta1.VirtualService
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1beta1.VirtualService) *v1beta1.VirtualService
	Update(ctx context.Context, VirtualService *v1beta1.VirtualService) *v1beta1.VirtualService
}

type VirtualService struct {
	cli *store.ClientConfig
	ns  string
}

func NewVirtualService(cli *store.ClientConfig, ns string) IVirtualService {
	return &VirtualService{cli: cli, ns: ns}
}

func (v *VirtualService) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageVirtualService {
	data, err := v.cli.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Lister().VirtualServices(v.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &form.PageVirtualService{}
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

func (v *VirtualService) Get(ctx context.Context, name string) *v1beta1.VirtualService {
	vs, err := v.cli.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Lister().VirtualServices(v.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return vs
}

func (v *VirtualService) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(v.cli.IstioClient.NetworkingV1beta1().VirtualServices(v.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func (v *VirtualService) Create(ctx context.Context, virtualService *v1beta1.VirtualService) *v1beta1.VirtualService {
	newVirtualService, err := v.cli.IstioClient.NetworkingV1beta1().VirtualServices(v.ns).Create(ctx, virtualService, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newVirtualService
}

func (v *VirtualService) Update(ctx context.Context, virtualService *v1beta1.VirtualService) *v1beta1.VirtualService {
	updateVirtualService, err := v.cli.IstioClient.NetworkingV1beta1().VirtualServices(v.ns).Update(ctx, virtualService, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateVirtualService
}

type VirtualServiceHandler struct {
	clusterName string
	client      *store.ClientConfig
}

func NewVirtualServiceHandler(client *store.ClientConfig, clusterName string) *VirtualServiceHandler {
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
		log.Logger.Error(err)
	}

	//按时间排序
	sort.SliceStable(VirtualServices, func(i, j int) bool {
		return VirtualServices[j].ObjectMeta.GetCreationTimestamp().Time.Before(VirtualServices[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	go wsstore.WsClientMap.SendClusterResource(v.clusterName, "virtualService", map[string]interface{}{
		"cluster": v.clusterName,
		"type":    "virtualService",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      VirtualServices,
		},
	})
}
