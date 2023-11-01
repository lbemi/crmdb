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

type SecretGetter interface {
	Secrets(namespace string) ISecret
}

type ISecret interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult
	Get(ctx context.Context, name string) *v1.Secret
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Secret) *v1.Secret
	Update(ctx context.Context, secret *v1.Secret) *v1.Secret
}

type Secret struct {
	client *store.ClientConfig
	ns     string
}

func NewSecret(client *store.ClientConfig, ns string) *Secret {
	return &Secret{client: client, ns: ns}
}

func (s *Secret) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data, err := s.client.SharedInformerFactory.Core().V1().Secrets().Lister().Secrets(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &form.PageResult{}
	var secretList = make([]*v1.Secret, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				secretList = append(secretList, item)
			}
		}
		data = secretList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				secretList = append(secretList, item)
			}
		}
		data = secretList
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

func (s *Secret) Get(ctx context.Context, name string) *v1.Secret {
	res, err := s.client.SharedInformerFactory.Core().V1().Secrets().Lister().Secrets(s.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *Secret) Delete(ctx context.Context, name string) {
	err := s.client.ClientSet.CoreV1().Secrets(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (s *Secret) Create(ctx context.Context, secret *v1.Secret) *v1.Secret {
	res, err := s.client.ClientSet.CoreV1().Secrets(s.ns).Create(ctx, secret, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (s *Secret) Update(ctx context.Context, secret *v1.Secret) *v1.Secret {
	res, err := s.client.ClientSet.CoreV1().Secrets(s.ns).Update(ctx, secret, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

type SecretHandle struct {
}

func NewSecretHandle() *SecretHandle {
	return &SecretHandle{}
}

func (s *SecretHandle) OnAdd(obj interface{}) {
	//TODO implement me
}

func (s *SecretHandle) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (s *SecretHandle) OnDelete(obj interface{}) {
	//TODO implement me
}
