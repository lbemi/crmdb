package global

import (
	"github.com/lbemi/lbemi/configs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Config
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
