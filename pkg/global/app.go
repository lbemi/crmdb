package global

import (
	"github.com/lbemi/lbemi/configs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Config
	Log         *zap.Logger
}

var App = new(Application)
