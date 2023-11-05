package services

import (
	"context"
	"github.com/lbemi/lbemi/apps/istio/api/vo"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"sort"
	"strings"

	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type GatewayGetter interface {
	Gateways(namespace string) IGateway
}

type IGateway interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *vo.PageGateway
	Get(ctx context.Context, name string) *v1beta1.Gateway
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1beta1.Gateway) *v1beta1.Gateway
	Update(ctx context.Context, Gateway *v1beta1.Gateway) *v1beta1.Gateway
}

type Gateways struct {
	cli *cache.ClientConfig
	ns  string
}

func NewGateway(cli *cache.ClientConfig, ns string) IGateway {
	return &Gateways{cli: cli, ns: ns}
}

func (v *Gateways) List(ctx context.Context, query *entity.PageParam, name string, label string) *vo.PageGateway {
	data, err := v.cli.IstioSharedInformerFactory.Networking().V1beta1().Gateways().Lister().Gateways(v.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &vo.PageGateway{}
	var GatewayList = make([]*v1beta1.Gateway, 0)

	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				GatewayList = append(GatewayList, item)
			}
		}
		data = GatewayList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				GatewayList = append(GatewayList, item)
			}
		}
		data = GatewayList
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

func (v *Gateways) Get(ctx context.Context, name string) *v1beta1.Gateway {
	vs, err := v.cli.IstioSharedInformerFactory.Networking().V1beta1().Gateways().Lister().Gateways(v.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return vs
}

func (v *Gateways) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(v.cli.IstioClient.NetworkingV1beta1().Gateways(v.ns).Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

func (v *Gateways) Create(ctx context.Context, Gateway *v1beta1.Gateway) *v1beta1.Gateway {
	newGateway, err := v.cli.IstioClient.NetworkingV1beta1().Gateways(v.ns).Create(ctx, Gateway, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newGateway
}

func (v *Gateways) Update(ctx context.Context, Gateway *v1beta1.Gateway) *v1beta1.Gateway {
	updateGateway, err := v.cli.IstioClient.NetworkingV1beta1().Gateways(v.ns).Update(ctx, Gateway, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return updateGateway
}

type GatewayHandler struct {
	clusterName string
	client      *cache.ClientConfig
}

func NewGatewayHandler(client *cache.ClientConfig, clusterName string) *GatewayHandler {
	return &GatewayHandler{client: client, clusterName: clusterName}
}

func (v *GatewayHandler) OnAdd(obj interface{}, isInInitialList bool) {
	v.notifyGateways(obj)
}

func (v *GatewayHandler) OnUpdate(oldObj, newObj interface{}) {
	v.notifyGateways(newObj)
}

func (v *GatewayHandler) OnDelete(obj interface{}) {
	v.notifyGateways(obj)
}

func (v *GatewayHandler) notifyGateways(obj interface{}) {
	namespace := obj.(*v1beta1.Gateway).Namespace
	Gateways, err := v.client.IstioSharedInformerFactory.Networking().V1beta1().Gateways().Lister().Gateways(namespace).List(labels.Everything())
	if err != nil {
		global.Logger.Error(err)
	}

	//按时间排序
	sort.SliceStable(Gateways, func(i, j int) bool {
		return Gateways[j].ObjectMeta.GetCreationTimestamp().Time.Before(Gateways[i].ObjectMeta.GetCreationTimestamp().Time)
	})

	go cache.WebsocketStore.SendClusterResource(v.clusterName, "Gateway", map[string]interface{}{
		"cluster": v.clusterName,
		"type":    "Gateway",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      Gateways,
		},
	})
}
