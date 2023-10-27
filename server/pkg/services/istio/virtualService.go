package istio

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type VirtualServiceImp interface {
	List(ctx context.Context) []*v1beta1.VirtualService
	Get(ctx context.Context, name string) *v1beta1.VirtualService
	Create(ctx context.Context, obj *v1beta1.VirtualService) *v1beta1.VirtualService
	Update(ctx context.Context, obj *v1beta1.VirtualService) *v1beta1.VirtualService
	Delete(ctx context.Context, name string)
}

type VirtualService struct {
	cli *store.ClientConfig
	ns  string
}

func (d *VirtualService) List(ctx context.Context) []*v1beta1.VirtualService {
	virtualServices, err := d.cli.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Lister().VirtualServices(d.ns).List(labels.Everything())
	//list, err := d.cli.IstioClient.NetworkingV1alpha3().VirtualServices(d.ns).List(ctx, metav1.ListOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	fmt.Println("virtualservice list: ===============", len(virtualServices))
	return virtualServices
}

func (d *VirtualService) Get(ctx context.Context, name string) *v1beta1.VirtualService {
	vs, err := d.cli.IstioSharedInformerFactory.Networking().V1beta1().VirtualServices().Lister().VirtualServices(d.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return vs
}

func (d *VirtualService) Create(ctx context.Context, obj *v1beta1.VirtualService) *v1beta1.VirtualService {
	newVirtualService, err := d.cli.IstioClient.NetworkingV1beta1().VirtualServices(d.ns).Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newVirtualService
}

func (d *VirtualService) Update(ctx context.Context, obj *v1beta1.VirtualService) *v1beta1.VirtualService {
	updateVirtualService, err := d.cli.IstioClient.NetworkingV1beta1().VirtualServices(d.ns).Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateVirtualService
}

func (d *VirtualService) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(d.cli.IstioClient.NetworkingV1beta1().VirtualServices(d.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func newVirtualService(cli *store.ClientConfig, namespace string) *VirtualService {
	return &VirtualService{cli: cli, ns: namespace}
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
