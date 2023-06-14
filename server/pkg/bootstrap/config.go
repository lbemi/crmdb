package bootstrap

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/model/config"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "./dev.yaml"
)

func InitializeConfig(configFile string) (appConfig *config.Config) {
	var config string
	if configFile == "" {
		config = defaultConfigFile
	} else {
		config = configFile
	}

	// ENV优先级大于指定--config
	if configEnv := os.Getenv("CONFIG"); configEnv != "" {
		config = configEnv
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read config failed. %v %s\n", err, config))
	}
	if err := v.Unmarshal(&appConfig); err != nil {
		panic(fmt.Sprintf("Unmarshal config failed. %v\n", err))
		return
	}
	return
}
