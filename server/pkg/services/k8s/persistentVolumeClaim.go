package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/restfulx"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type PersistentVolumeClaimImp interface {
	List(ctx context.Context) []*v1.PersistentVolumeClaim
	Get(ctx context.Context, name string) *v1.PersistentVolumeClaim
	Delete(ctx context.Context, name string)
	Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
	Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
}

type persistentVolumeClaim struct {
	client *store.ClientConfig
	ns     string
}

func (s *persistentVolumeClaim) List(ctx context.Context) []*v1.PersistentVolumeClaim {
	nodeList, err := s.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(s.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return nodeList
}

func (s *persistentVolumeClaim) Get(ctx context.Context, name string) *v1.PersistentVolumeClaim {
	res, err := s.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(s.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *persistentVolumeClaim) Delete(ctx context.Context, name string) {
	err := s.client.ClientSet.CoreV1().PersistentVolumeClaims(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (s *persistentVolumeClaim) Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	res, err := s.client.ClientSet.CoreV1().PersistentVolumeClaims(s.ns).Create(ctx, pvc, metav1.CreateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func (s *persistentVolumeClaim) Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	res, err := s.client.ClientSet.CoreV1().PersistentVolumeClaims(s.ns).Update(ctx, pvc, metav1.UpdateOptions{})
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return res
}

func newPersistentVolumeClaim(client *store.ClientConfig, namespace string) *persistentVolumeClaim {
	return &persistentVolumeClaim{client: client, ns: namespace}
}

type PersistentVolumeClaimHandler struct {
	client      *store.ClientConfig
	clusterName string
}

func NewPersistentVolumeClaimHandler(client *store.ClientConfig, clusterName string) *PersistentVolumeClaimHandler {
	return &PersistentVolumeClaimHandler{client: client, clusterName: clusterName}
}

func (p *PersistentVolumeClaimHandler) OnAdd(obj interface{}) {
	//TODO implement me
}

func (p *PersistentVolumeClaimHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
}

func (p *PersistentVolumeClaimHandler) OnDelete(obj interface{}) {
	//TODO implement me
}
