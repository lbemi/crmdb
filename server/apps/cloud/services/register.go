package services

import (
	"github.com/lbemi/lbemi/apps/cloud/api/form"
	"github.com/lbemi/lbemi/apps/cloud/entity"
	istioHandler "github.com/lbemi/lbemi/apps/istio/services"
	"github.com/lbemi/lbemi/apps/kubernetes/services"
	tektonGetter "github.com/lbemi/lbemi/apps/tekton/service"
	"github.com/lbemi/lbemi/pkg/cache"
	"gorm.io/gorm"
)

type ClusterGetter interface {
	Cluster(clusterName string) ICluster
}

type ICluster interface {
	Create(config *form.ClusterReq)
	Delete(id uint64)
	Update(id uint64, config *entity.Cluster)
	Get(id uint64) *entity.Cluster
	List() *[]entity.Cluster
	GetByName(name string) *entity.Cluster
	ChangeStatus(id uint64, status bool)
	CheckHealth() bool
	GenerateClient(name, config string) (*cache.ClientConfig, *entity.Cluster, error)

	istioHandler.IstioGetter
	services.K8SGetter
	tektonGetter.TektonGetter
}

type Cluster struct {
	clusterName string
	db          *gorm.DB
	store       *cache.ClientStore
}

func (c *Cluster) Istio() istioHandler.IstioInterface {
	return istioHandler.NewIstio(c.clusterName, c.store)
}
func (c *Cluster) K8S() services.K8SInterface {
	return services.NewK8S(c.clusterName, c.store)
}
func (c *Cluster) Tekton() tektonGetter.TektonInterface {
	return tektonGetter.NewTekton(c.clusterName, c.store)
}
