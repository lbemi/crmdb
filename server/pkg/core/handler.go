package core

import (
	"github.com/casbin/casbin/v2"
	r "github.com/go-redis/redis"
	services2 "github.com/lbemi/lbemi/apps/asset/services"
	services3 "github.com/lbemi/lbemi/apps/cloud/services"
	services4 "github.com/lbemi/lbemi/apps/log/services"
	"github.com/lbemi/lbemi/apps/system/services"
	"github.com/lbemi/lbemi/pkg/bootstrap/policy"
	"github.com/lbemi/lbemi/pkg/bootstrap/redis"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/config"
	"gorm.io/gorm"
)

type Getter interface {
	redis.RedisGetter
	policy.PolicyGetter
	services.UserGetter
	services.RoleGetter
	services.MenuGetter
	services2.HostGetter
	services2.AccountGetter
	services2.TerminalGetter
	services2.GroupGetter
	services2.WsGetter
	services3.ClusterGetter
	services4.LoginLogGetter
	services4.OperatorLogGetter
	services2.ResourceAccountGetter
}

type Handler struct {
	Config      *config.Config
	RedisCli    *r.Client
	DB          *gorm.DB
	Enforcer    *casbin.SyncedEnforcer
	ClientStore *cache.ClientMap
}

func NewHandler(redisCli *r.Client, db *gorm.DB, enforcer *casbin.SyncedEnforcer, clientStore *cache.ClientMap) Getter {
	return &Handler{
		RedisCli:    redisCli,
		DB:          db,
		Enforcer:    enforcer,
		ClientStore: clientStore,
	}
}

func (c *Handler) Redis() redis.IRedis {
	return redis.NewRedis(c.RedisCli)
}

func (c *Handler) Policy() policy.IPolicy {
	return policy.NewPolicy(c.DB, c.Enforcer)
}

func (c *Handler) User() services.IUSer {
	return services.NewUser(c.DB, c.Policy(), c.Menu(), c.Login())
}

func (c *Handler) Role() services.IRole {
	return services.NewRole(c.DB, c.Policy(), c.Menu())
}

func (c *Handler) Menu() services.IMenu {
	return services.NewMenu(c.DB, c.Policy())
}

func (c *Handler) Host() services2.IHost {
	return services2.NewHost(c.DB, c.ResourceBindAccount(), c.Account())
}

func (c *Handler) Account() services2.IAccount {
	return services2.NewAccount(c.DB)
}

func (c *Handler) Group() services2.IGroup {
	return services2.NewGroup(c.DB)
}

func (c *Handler) Terminal() services2.ITerminal {
	return services2.NewTerminal(c.Host(), c.Account())
}

func (c *Handler) Ws() services2.IWs {
	return services2.NewWs()
}

func (c *Handler) Cluster(clusterName string) services3.ICluster {
	return services3.NewCluster(c.DB, c.ClientStore, clusterName)
}

func (c *Handler) Login() services4.ILoginLog {
	return services4.NewLogin(c.DB)
}

func (c *Handler) Operator() services4.IOperatorLog {
	return services4.NewOperator(c.DB)
}

func (c *Handler) ResourceBindAccount() services2.IResourceBindAccount {
	return services2.NewResourceBindAccount(c.DB)
}
