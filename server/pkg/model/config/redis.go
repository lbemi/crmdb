package config

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}
