package option

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/config"
	"github.com/lbemi/lbemi/pkg/services"
	"gorm.io/gorm"
)

type Options struct {
	Factory   services.IDbFactory
	Config    *config.Config
	DB        *gorm.DB
	Redis     *redis.Client
	Enforcer  *casbin.Enforcer
	GinEngine *gin.Engine
}

func (o *Options) Load() {
	// 加载配置文件
	o.Config = bootstrap.InitializeConfig()
	// 初始化日志
	log.Register(&o.Config.Log)
	// 初始化数据库
	o.DB = bootstrap.InitializeDB(o.Config)
	// 初始化redis
	o.Redis = bootstrap.InitializeRedis(o.Config.Redis)
	// 注册校验器
	bootstrap.InitializeValidator()
	// 初始化ginEngine
	o.GinEngine = gin.New()
	// 初始化casbin enforcer
	o.Enforcer = bootstrap.InitPolicyEnforcer(o.DB)
	// 初始化dbFactory
	o.Factory = services.NewDbFactory(o.DB, o.Enforcer)
}

func NewOptions() *Options {
	return &Options{}
}
