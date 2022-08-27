package configs

type Config struct {
	App      `mapstructure:"app" `
	Log      `mapstructure:"log"`
	Database `mapstructure:"database"`
	Jwt      `mapstructure:"jwt"`
}
