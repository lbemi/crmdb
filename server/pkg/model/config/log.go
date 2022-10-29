package config

type Log struct {
	Level     string `mapstructure:"level"`
	RootDir   string `mapstructure:"root_dir"`
	FileName  string `mapstructure:"file_name"`
	Format    string `mapstructure:"format"`
	ShowLine  bool   `mapstructure:"show_line"`
	MaxBackup int    `mapstructure:"max_backup"`
	MaxSize   int    `mapstructure:"max_size"`
	MaxAge    int    `mapstructure:"max_age" `
	Compress  bool   `mapstructure:"compress"`
	IsFile    bool   `mapstructure:"is_file"`
}
