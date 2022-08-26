package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/spf13/viper"
	"os"
)

const (
	defaultConfigFile = "../../config.yaml"
)

func InitializeConfig() *viper.Viper {
	config := defaultConfigFile
	if configEnv := os.Getenv("CONFIG"); configEnv != "" {
		config = configEnv
	}
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read config failed. %v\n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Print("config file changed:", in.Name)
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
