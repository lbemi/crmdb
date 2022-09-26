package common

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lbemi/lbemi/pkg/model/configs"
	"github.com/spf13/viper"
)

func InitConfig(configPath string) (conf *configs.Config, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
		return nil, err
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
		return nil, err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&conf)
		if err != nil {
			panic(err)
		}
	})
	return conf, nil
}
