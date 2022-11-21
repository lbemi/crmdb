package bootstrap

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/model/config"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "./config.yaml"
)

func InitializeConfig() (appConfig *config.Config) {
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
	//v.WatchConfig()
	//v.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Print("config file changed:", in.Name)
	//	if err := v.Unmarshal(&global.App.Config); err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//})
	if err := v.Unmarshal(&appConfig); err != nil {
		panic(fmt.Sprintf("Unmarshal config failed. %v\n", err))
		return
	}
	return
}
