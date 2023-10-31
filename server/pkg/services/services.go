package services

import (
	"github.com/lbemi/lbemi/pkg/common/store"
	client "github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/services/asset"
	"github.com/lbemi/lbemi/pkg/services/auth"
	"github.com/lbemi/lbemi/pkg/services/cloud"
	"github.com/lbemi/lbemi/pkg/services/logsys"
	"github.com/lbemi/lbemi/pkg/services/sys"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

type Interface interface {
	Authentication() auth.IAuthentication
	User() sys.IUSer
	Role() sys.IRole
	Menu() sys.IMenu
	Host() asset.IHost
	Account() asset.IAccount
	Group() asset.IGroup
	Cluster() cloud.ICluster
	Log() logsys.ILoginLog
	Operator() logsys.IOperatorLog
	ResourceBindAccount() asset.IResourceAccount
}

type DbFactory struct {
	db      *gorm.DB
	enforce *casbin.SyncedEnforcer
	client  *client.KubernetesClient
	store   *store.ClientMap
}

func (f *DbFactory) Authentication() auth.IAuthentication {
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
func (f *DbFactory) Account() asset.IAccount {
	return asset.NewAccount(f.db)
}

func (f *DbFactory) Group() asset.IGroup {
	return asset.NewGroup(f.db)
}

func (f *DbFactory) Cluster() cloud.ICluster {
	return cloud.NewCluster(f.db, f.store)
}

func (f *DbFactory) Log() logsys.ILoginLog {
	return logsys.NewLoginLog(f.db)
}
func (f *DbFactory) Operator() logsys.IOperatorLog {
	return logsys.NewOperatorLog(f.db)
}

func (f *DbFactory) ResourceBindAccount() asset.IResourceAccount {
	return asset.NewResourceBindAccount(f.db)
}

func NewDbFactory(db *gorm.DB, enforcer *casbin.SyncedEnforcer, store *store.ClientMap) Interface {
	return &DbFactory{
		db:      db,
		enforce: enforcer,
		store:   store,
	}
}
