package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/restfulx"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sort"
)

type ServiceImp interface {
	List(ctx context.Context) []*v1.Service
	Get(ctx context.Context, name string) *v1.Service
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, node *v1.Service) *v1.Service
	Update(ctx context.Context, service *v1.Service) *v1.Service
	ListWorkLoad(ctx context.Context, name string) *types.ServiceWorkLoad
}

type service struct {
	client *store.ClientConfig
	ns     string
}

func (s *service) List(ctx context.Context) []*v1.Service {
	serviceList, err := s.client.SharedInformerFactory.Core().V1().Services().Lister().Services(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	//按时间排序
	sort.Slice(serviceList, func(i, j int) bool {
		return serviceList[j].ObjectMeta.GetCreationTimestamp().Time.Before(serviceList[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	return serviceList
}
func (s *service) ListWorkLoad(ctx context.Context, name string) *types.ServiceWorkLoad {
	workLoad := &types.ServiceWorkLoad{
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

func (s *service) Get(ctx context.Context, name string) *v1.Service {
	res, err := s.client.SharedInformerFactory.Core().V1().Services().Lister().Services(s.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *service) Delete(ctx context.Context, name string) {
	err := s.client.ClientSet.CoreV1().Services(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (s *service) Create(ctx context.Context, service *v1.Service) *v1.Service {
	res, err := s.client.ClientSet.CoreV1().Services(s.ns).Create(ctx, service, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (s *service) Update(ctx context.Context, service *v1.Service) *v1.Service {
	res, err := s.client.ClientSet.CoreV1().Services(s.ns).Update(ctx, service, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func newService(client *store.ClientConfig, namespace string) *service {
	return &service{client: client, ns: namespace}
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
