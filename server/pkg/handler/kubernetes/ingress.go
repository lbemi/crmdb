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
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type IngressesGetter interface {
	Ingresses(namespace string) IIngresses
}

type IIngresses interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.Ingress
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Ingress) *v1.Ingress
	Update(ctx context.Context, ingresses *v1.Ingress) *v1.Ingress
}

type ingresses struct {
	client    *store.ClientConfig
	namespace string
}

func NewIngresses(client *store.ClientConfig, namespace string) *ingresses {
	return &ingresses{client: client, namespace: namespace}
}

func (i *ingresses) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data, err := i.client.SharedInformerFactory.Networking().V1().Ingresses().Lister().Ingresses(i.namespace).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &form.PageResult{}
	var ingressList = make([]*v1.Ingress, 0)

	for _, item := range data {
		if (name == "" || strings.Contains(item.Name, name)) && (label == "" || strings.Contains(labels.FormatLabels(item.Labels), label)) {
			ingressList = append(ingressList, item)
		}
	}
	sort.Slice(ingressList, func(i, j int) bool {
		return ingressList[j].ObjectMeta.CreationTimestamp.Time.Before(ingressList[i].ObjectMeta.CreationTimestamp.Time)
	})

	total := len(ingressList)
	if query.Limit == 0 && query.Page == 0 {
		res.Data = ingressList
	} else {
		if total <= query.Limit {
			res.Data = ingressList
		} else if query.Page*query.Limit >= total {
			res.Data = ingressList[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = ingressList[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}

	res.Total = int64(total)
	return res
}

func (i *ingresses) Get(ctx context.Context, name string) *v1.Ingress {
	res, err := i.client.SharedInformerFactory.Networking().V1().Ingresses().Lister().Ingresses(i.namespace).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (i *ingresses) Delete(ctx context.Context, name string) {
	err := i.client.ClientSet.NetworkingV1().Ingresses(i.namespace).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (i *ingresses) Create(ctx context.Context, ingresses *v1.Ingress) *v1.Ingress {
	res, err := i.client.ClientSet.NetworkingV1().Ingresses(i.namespace).Create(ctx, ingresses, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (i *ingresses) Update(ctx context.Context, ingresses *v1.Ingress) *v1.Ingress {
	res, err := i.client.ClientSet.NetworkingV1().Ingresses(i.namespace).Update(ctx, ingresses, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

type IngressHandle struct{}

func NewIngressHandle() *IngressHandle {
	return &IngressHandle{}
}

func (i *IngressHandle) OnAdd(obj interface{}) {
	//TODO implement me
}

func (i *IngressHandle) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (i *IngressHandle) OnDelete(obj interface{}) {
	//TODO implement me
}
