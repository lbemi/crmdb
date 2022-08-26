package configs

type Config struct {
	App `mapstructure:"app" `
	Log `mapstructure:"log"`
}
