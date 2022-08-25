package global

import (
	"github.com/lbemi/lbemi/configs"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Config
}

var App = new(Application)
