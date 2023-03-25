package services

import (
	"github.com/casbin/casbin/v2"
	"github.com/lbemi/lbemi/pkg/common/store"
	client "github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/services/asset"
	"github.com/lbemi/lbemi/pkg/services/auth"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	"github.com/lbemi/lbemi/pkg/services/sys"
	"gorm.io/gorm"
)

type FactoryImp interface {
	Authentication() auth.AuthenticationInterface
	User() sys.IUSer
	Role() sys.IRole
	Menu() sys.IMenu
	Host() asset.IHost
	Cluster() cloud.ICluster
}

type DbFactory struct {
	db      *gorm.DB
	enforce *casbin.Enforcer
	client  *client.KubernetesClient
	store   *store.ClientStore
}

func (f *DbFactory) Authentication() auth.AuthenticationInterface {
	return auth.NewAuthentication(f.db, f.enforce)
}

func (f *DbFactory) User() sys.IUSer {
	return sys.NewUser(f.db)
}

func (f *DbFactory) Role() sys.IRole {
	return sys.NewRole(f.db)
}

func (f *DbFactory) Menu() sys.IMenu {
	return sys.NewMenu(f.db)
}

func (f *DbFactory) Host() asset.IHost {
	return asset.NewHost(f.db)
}

func (f *DbFactory) Cluster() cloud.ICluster {
	return cloud.NewCluster(f.db, f.store)
}

func NewDbFactory(db *gorm.DB, enforcer *casbin.Enforcer, store *store.ClientStore) FactoryImp {
	return &DbFactory{
		db:      db,
		enforce: enforcer,
		store:   store,
	}
}
