package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis"
	"github.com/lbemi/lbemi/pkg/model/configs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Config
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
	Enforcer    *casbin.Enforcer
}

var App = new(Application)
