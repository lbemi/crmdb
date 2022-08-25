package configs

//App 项目配置
type App struct {
	Addr    string `yaml:"addr"`
	Port    string `yaml:"port"`
	Env     string `yaml:"env"`
	AppName string `yaml:"app_name"`
	AppUrl  string `yaml:"app_url"`
}
