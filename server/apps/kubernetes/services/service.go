package services

import (
	"context"
	entity2 "github.com/lbemi/lbemi/apps/kubernetes/entity"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ServiceGetter interface {
	Service(namespace string) IService
}

type IService interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *v1.Service
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Service) *v1.Service
	Update(ctx context.Context, service *v1.Service) *v1.Service
	ListWorkLoad(ctx context.Context, name string) *entity2.ServiceWorkLoad
}

type Service struct {
	client *cache.ClientConfig
	ns     string
}

func NewService(client *cache.ClientConfig, namespace string) *Service {
	return &Service{client: client, ns: namespace}
}

func (s *Service) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := s.client.SharedInformerFactory.Core().V1().Services().Lister().Services(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var serviceList = make([]*v1.Service, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				serviceList = append(serviceList, item)
			}
		}
		data = serviceList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				serviceList = append(serviceList, item)
			}
		}
		data = serviceList
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

func (s *Service) ListWorkLoad(ctx context.Context, name string) *entity2.ServiceWorkLoad {
	workLoad := &entity2.ServiceWorkLoad{
		Deployments:  make([]*appsv1.Deployment, 0),
		StatefulSets: make([]*appsv1.StatefulSet, 0),
		DaemonSets:   make([]*appsv1.DaemonSet, 0),
		Events:       make([]*v1.Event, 0),
	}
	svc := s.Get(ctx, name)
	selector := labels.SelectorFromSet(svc.Spec.Selector)

	if selector.Empty() {
		return workLoad
	}
	eventList, err := s.client.SharedInformerFactory.Core().V1().Events().Lister().Events(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	for _, item := range eventList {
		if item.InvolvedObject.Kind == "Service" && item.InvolvedObject.Name == name {
			workLoad.Events = append(workLoad.Events, item)
		}
	}
	endpoints, err := s.client.ClientSet.CoreV1().Endpoints(s.ns).Get(ctx, name, metav1.GetOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	deployments, err := s.client.SharedInformerFactory.Apps().V1().Deployments().Lister().List(selector)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	daemonSets, err := s.client.SharedInformerFactory.Apps().V1().DaemonSets().Lister().List(selector)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	statefulSets, err := s.client.SharedInformerFactory.Apps().V1().StatefulSets().Lister().List(selector)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)

	workLoad.Deployments = deployments
	workLoad.StatefulSets = statefulSets
	workLoad.DaemonSets = daemonSets
	workLoad.EndPoints = endpoints

	return workLoad
}

func (s *Service) Get(ctx context.Context, name string) *v1.Service {
	res, err := s.client.SharedInformerFactory.Core().V1().Services().Lister().Services(s.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *Service) Delete(ctx context.Context, name string) {
	err := s.client.ClientSet.CoreV1().Services(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (s *Service) Create(ctx context.Context, service *v1.Service) *v1.Service {
	res, err := s.client.ClientSet.CoreV1().Services(s.ns).Create(ctx, service, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (s *Service) Update(ctx context.Context, service *v1.Service) *v1.Service {
	res, err := s.client.ClientSet.CoreV1().Services(s.ns).Update(ctx, service, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

type ServiceHandle struct{}

func NewServiceHandle() *ServiceHandle {
	return &ServiceHandle{}
}

func (s *ServiceHandle) OnAdd(obj interface{}) {
	//TODO implement me
}

func (s *ServiceHandle) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (s *ServiceHandle) OnDelete(obj interface{}) {
	//TODO implement me
}
