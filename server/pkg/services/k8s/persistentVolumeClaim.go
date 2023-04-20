package k8s

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type PersistentVolumeClaimImp interface {
	List(ctx context.Context) ([]*v1.PersistentVolumeClaim, error)
	Get(ctx context.Context, name string) (*v1.PersistentVolumeClaim, error)
	Delete(ctx context.Context, name string) error
	Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
	Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
}

type persistentVolumeClaim struct {
	client *store.Clients
	ns     string
}

func (s *persistentVolumeClaim) List(ctx context.Context) ([]*v1.PersistentVolumeClaim, error) {
	nodeList, err := s.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(s.ns).List(labels.Everything())
	if err != nil {
		log.Logger.Error(err)
	}
	return nodeList, err
}

func (s *persistentVolumeClaim) Get(ctx context.Context, name string) (*v1.PersistentVolumeClaim, error) {
	res, err := s.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(s.ns).Get(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *persistentVolumeClaim) Delete(ctx context.Context, name string) error {
	err := s.client.ClientSet.CoreV1().PersistentVolumeClaims(s.ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *persistentVolumeClaim) Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	res, err := s.client.ClientSet.CoreV1().PersistentVolumeClaims(s.ns).Create(ctx, pvc, metav1.CreateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func (s *persistentVolumeClaim) Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	res, err := s.client.ClientSet.CoreV1().PersistentVolumeClaims(s.ns).Update(ctx, pvc, metav1.UpdateOptions{})
	if err != nil {
		log.Logger.Error(err)
	}
	return res, err
}

func newPersistentVolumeClaim(client *store.Clients, namespace string) *persistentVolumeClaim {
	return &persistentVolumeClaim{client: client, ns: namespace}
}

type PersistentVolumeClaimHandler struct {
	client      *store.Clients
	clusterName string
}

func NewPersistentVolumeClaimHandler(client *store.Clients, clusterName string) *PersistentVolumeClaimHandler {
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
