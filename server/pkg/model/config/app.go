package config

// App 项目配置
type App struct {
	Addr    string `mapstructure:"addr"`
	Port    string `mapstructure:"port"`
	Env     string `mapstructure:"env"`
	AppName string `mapstructure:"app_name"`
	AppUrl  string `mapstructure:"app_url"`
}
