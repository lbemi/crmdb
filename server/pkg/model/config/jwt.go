package config

type Jwt struct {
	Key string `mapstructure:"key" json:"key"`
	TTL int64  `mapstructure:"ttl"`
}
