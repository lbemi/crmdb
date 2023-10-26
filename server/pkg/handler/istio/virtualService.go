package istio

import (
	"context"
	"k8s.io/apimachinery/pkg/labels"
	"strings"

	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services/istio"

	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

type VirtualServiceGetter interface {
	VirtualServices(namespace string) IVirtualService
}

type IVirtualService interface {
	List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageVirtualService
	Get(ctx context.Context, name string) *v1alpha3.VirtualService
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1alpha3.VirtualService) *v1alpha3.VirtualService
	Update(ctx context.Context, VirtualService *v1alpha3.VirtualService) *v1alpha3.VirtualService
}

type VirtualService struct {
	istio *istio.Factory
}

func NewVirtualService(istio *istio.Factory) *VirtualService {
	return &VirtualService{istio: istio}
}

func (s *VirtualService) List(ctx context.Context, query *model.PageParam, name string, label string) *form.PageVirtualService {
	data := s.istio.VirtualService().List(ctx)
	res := &form.PageVirtualService{}
	var VirtualServiceList = make([]*v1alpha3.VirtualService, 0)

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

func (s *VirtualService) Get(ctx context.Context, name string) *v1alpha3.VirtualService {
	return s.istio.VirtualService().Get(ctx, name)
}

func (s *VirtualService) Delete(ctx context.Context, name string) {
	s.istio.VirtualService().Delete(ctx, name)
}

func (s *VirtualService) Create(ctx context.Context, VirtualService *v1alpha3.VirtualService) *v1alpha3.VirtualService {
	return s.istio.VirtualService().Create(ctx, VirtualService)
}

func (s *VirtualService) Update(ctx context.Context, VirtualService *v1alpha3.VirtualService) *v1alpha3.VirtualService {
	return s.istio.VirtualService().Update(ctx, VirtualService)
}
