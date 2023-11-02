package core

import (
	"github.com/casbin/casbin/v2"
	r "github.com/go-redis/redis"
	asset "github.com/lbemi/lbemi/apps/asset/services"
	cloud "github.com/lbemi/lbemi/apps/cloud/services"
	logsys "github.com/lbemi/lbemi/apps/log/services"
	"github.com/lbemi/lbemi/apps/system/services"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/common/commService"
	"github.com/lbemi/lbemi/pkg/config"
	"gorm.io/gorm"
)

var V1 getter

func Register(options *option.Options) {
	V1 = newService(
		options.Redis,
		options.DB,
		options.Enforcer,
		options.ClientStore,
	)
}

// 聚合service,实现链式掉用
type getter interface {
	commService.RedisGetter
	commService.PolicyGetter
	services.UserGetter
	services.RoleGetter
	services.MenuGetter
	asset.HostGetter
	asset.AccountGetter
	asset.TerminalGetter
	asset.GroupGetter
	asset.WsGetter
	cloud.ClusterGetter
	logsys.LoginLogGetter
	logsys.OperatorLogGetter
	asset.ResourceAccountGetter
}

type service struct {
	Config      *config.Config
	RedisCli    *r.Client
	DB          *gorm.DB
	Enforcer    *casbin.SyncedEnforcer
	ClientStore *cache.ClientStore
}

func newService(redisCli *r.Client, db *gorm.DB, enforcer *casbin.SyncedEnforcer, clientStore *cache.ClientStore) getter {
	return &service{
		RedisCli:    redisCli,
		DB:          db,
		Enforcer:    enforcer,
		ClientStore: clientStore,
	}
}

func (c *service) Redis() commService.IRedis {
	return commService.NewRedis(c.RedisCli)
}

func (c *service) Policy() commService.IPolicy {
	return commService.NewPolicy(c.DB, c.Enforcer)
}

func (c *service) User() services.IUSer {
	return services.NewUser(c.DB, c.Policy(), c.Menu(), c.Login())
}

func (c *service) Role() services.IRole {
	return services.NewRole(c.DB, c.Policy(), c.Menu())
}

func (c *service) Menu() services.IMenu {
	return services.NewMenu(c.DB, c.Policy())
}

func (c *service) Host() asset.IHost {
	return asset.NewHost(c.DB, c.ResourceBindAccount(), c.Account())
}

func (c *service) Account() asset.IAccount {
	return asset.NewAccount(c.DB)
}

func (c *service) Group() asset.IGroup {
	return asset.NewGroup(c.DB)
}

func (c *service) Terminal() asset.ITerminal {
	return asset.NewTerminal(c.Host(), c.Account())
}

func (c *service) Ws() asset.IWs {
	return asset.NewWs()
}

func (c *service) Cluster(clusterName string) cloud.ICluster {
	return cloud.NewCluster(c.DB, c.ClientStore, clusterName)
}

func (c *service) Login() logsys.ILoginLog {
	return logsys.NewLogin(c.DB)
}

func (c *service) Operator() logsys.IOperatorLog {
	return logsys.NewOperator(c.DB)
}

func (c *service) ResourceBindAccount() asset.IResourceBindAccount {
	return asset.NewResourceBindAccount(c.DB)
}
