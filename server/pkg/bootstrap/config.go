package bootstrap

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/config"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "./config.yaml"
)

func InitializeConfig(configFile string) (appConfig *config.Config) {
	var configName string
	if configFile == "" {
		configName = defaultConfigFile
	} else {
		configName = configFile
	}

	// ENV优先级大于指定--configName
	if configEnv := os.Getenv("CONFIG"); configEnv != "" {
		configName = configEnv
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(configName)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read configName failed. %v %s\n", err, configName))
	}
	if err := v.Unmarshal(&appConfig); err != nil {
		panic(fmt.Sprintf("Unmarshal configName failed. %v\n", err))
		return
	}
	return
}
