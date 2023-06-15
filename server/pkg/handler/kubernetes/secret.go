package kubernetes

import (
	"context"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/k8s"

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

type secret struct {
	k8s *k8s.Factory
}

func NewSecret(k8s *k8s.Factory) *secret {
	return &secret{k8s: k8s}
}

func (s *secret) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageResult {
	data := s.k8s.Secret().List(ctx)
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

func (s *secret) Get(ctx context.Context, name string) *v1.Secret {
	return s.k8s.Secret().Get(ctx, name)
}

func (s *secret) Delete(ctx context.Context, name string) {
	s.k8s.Secret().Delete(ctx, name)
}

func (s *secret) Create(ctx context.Context, secret *v1.Secret) *v1.Secret {
	return s.k8s.Secret().Create(ctx, secret)
}

func (s *secret) Update(ctx context.Context, secret *v1.Secret) *v1.Secret {
	return s.k8s.Secret().Update(ctx, secret)
}
