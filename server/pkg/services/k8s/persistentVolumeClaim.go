package k8s

//
//import (
//	"context"
//	"sort"
//
//	"github.com/lbemi/lbemi/pkg/common/store"
//	"github.com/lbemi/lbemi/pkg/restfulx"
//
//	v1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/labels"
//)
//
//type PersistentVolumeClaimImp interface {
//	List(ctx context.Context) []*v1.PersistentVolumeClaim
//	Get(ctx context.Context, name string) *v1.PersistentVolumeClaim
//	Delete(ctx context.Context, name string)
//	Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
//	Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim
//}
//
//type persistentVolumeClaim struct {
//	client *store.ClientConfig
//	ns     string
//}
//
//func (p *persistentVolumeClaim) List(ctx context.Context) []*v1.PersistentVolumeClaim {
//	list, err := p.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(p.ns).List(labels.Everything())
//	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
//	sort.Slice(list, func(i, j int) bool {
//		return list[j].ObjectMeta.CreationTimestamp.Time.Before(list[i].ObjectMeta.CreationTimestamp.Time)
//	})
//	return list
//}
//
//func (p *persistentVolumeClaim) Get(ctx context.Context, name string) *v1.PersistentVolumeClaim {
//	res, err := p.client.SharedInformerFactory.Core().V1().PersistentVolumeClaims().Lister().PersistentVolumeClaims(p.ns).Get(name)
//	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
//	return res
//}
//
//func (p *persistentVolumeClaim) Delete(ctx context.Context, name string) {
//	err := p.client.ClientSet.CoreV1().PersistentVolumeClaims(p.ns).Delete(ctx, name, metav1.DeleteOptions{})
//	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
//}
//
//func (p *persistentVolumeClaim) Create(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
//	res, err := p.client.ClientSet.CoreV1().PersistentVolumeClaims(p.ns).Create(ctx, pvc, metav1.CreateOptions{})
//	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
//	return res
//}
//
//func (p *persistentVolumeClaim) Update(ctx context.Context, pvc *v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
//	res, err := p.client.ClientSet.CoreV1().PersistentVolumeClaims(p.ns).Update(ctx, pvc, metav1.UpdateOptions{})
//	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
//	return res
//}
//
//func newPersistentVolumeClaim(client *store.ClientConfig, namespace string) *persistentVolumeClaim {
//	return &persistentVolumeClaim{client: client, ns: namespace}
//}
//
//type PersistentVolumeClaimHandler struct {
//	client      *store.ClientConfig
//	clusterName string
//}
//
//func NewPersistentVolumeClaimHandler(client *store.ClientConfig, clusterName string) *PersistentVolumeClaimHandler {
//	return &PersistentVolumeClaimHandler{client: client, clusterName: clusterName}
//}
//
//func (p *PersistentVolumeClaimHandler) OnAdd(obj interface{}) {
//	//TODO implement me
//}
//
//func (p *PersistentVolumeClaimHandler) OnUpdate(oldObj, newObj interface{}) {
//	//TODO implement me
//}
//
//func (p *PersistentVolumeClaimHandler) OnDelete(obj interface{}) {
//	//TODO implement me
//}
