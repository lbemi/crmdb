package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type NamespaceGetter interface {
	Namespaces() INamespace
}

type INamespace interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.Namespace
	Create(ctx context.Context, obj *v1.Namespace) *v1.Namespace
	Update(ctx context.Context, obj *v1.Namespace) *v1.Namespace
	Delete(ctx context.Context, name string)
}

type namespace struct {
	cli *cache.ClientConfig
}

func NewNamespace(cli *cache.ClientConfig) *namespace {
	return &namespace{cli: cli}
}

func (n *namespace) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := n.cli.SharedInformerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	namespaceList := make([]*v1.Namespace, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				namespaceList = append(namespaceList, item)
			}
		}
		data = namespaceList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				namespaceList = append(namespaceList, item)
			}
		}
		data = namespaceList
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

func (n *namespace) Get(ctx context.Context, name string) *v1.Namespace {
	dep, err := n.cli.SharedInformerFactory.Core().V1().Namespaces().Lister().Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return dep
}

func (n *namespace) Create(ctx context.Context, obj *v1.Namespace) *v1.Namespace {
	newNamespace, err := n.cli.ClientSet.CoreV1().Namespaces().Create(ctx, obj, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newNamespace
}

func (n *namespace) Update(ctx context.Context, obj *v1.Namespace) *v1.Namespace {
	newNamespace, err := n.cli.ClientSet.CoreV1().Namespaces().Update(ctx, obj, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return newNamespace
}

func (n *namespace) Delete(ctx context.Context, name string) {
	restfulx.ErrNotNilDebug(n.cli.ClientSet.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{}), restfulx.OperatorErr)
}

type NameSpaceHandler struct {
	cli *cache.ClientConfig
	ns  string
}

func NewNameSpaceHandler(cli *cache.ClientConfig, ns string) *NameSpaceHandler {
	return &NameSpaceHandler{cli: cli, ns: ns}
}

func (n *NameSpaceHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (n *NameSpaceHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (n *NameSpaceHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
