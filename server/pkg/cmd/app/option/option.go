package option

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/config"
	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
)

type Options struct {
	Config      *config.Config
	DB          *gorm.DB
	Redis       *redis.Client
	Enforcer    *casbin.SyncedEnforcer
	GinEngine   *gin.Engine
	ClientStore *cache.ClientMap
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) WithConfig(configFile string) *Options {
	// 加载配置文件
	o.Config = bootstrap.InitializeConfig(configFile)
	return o
}
func (o *Options) WithLog() *Options {
	// 加载配置文件
	log.Register(&o.Config.Log)
	return o
}
func (o *Options) Complete() *Options {
	// 加载配置文件
	//o.Config = bootstrap.InitializeConfig()
	// 初始化日志
	log.Register(&o.Config.Log)
	// 初始化数据库
	o.DB = bootstrap.InitializeDB(o.Config)
	// 初始化redis
	o.Redis = bootstrap.InitializeRedis(o.Config.Redis)
	// 注册校验器
	bootstrap.InitializeValidator()
	// 初始化ginEngine
	//o.GinEngine = gin.New()
	// 初始化casbin enforcer
	o.Enforcer = bootstrap.InitPolicyEnforcer(o.DB)
	// 初始化client store
	clientStore := cache.NewClientStore()
	o.ClientStore = clientStore
	return o
}
