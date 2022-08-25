package configs

type Log struct {
	Level     string `yaml:"level"`
	RootDir   string `yaml:"root_dir"`
	FileName  string `yaml:"file_name"`
	Format    string `yaml:"format"`
	ShowLine  bool   `yaml:"show_line"`
	MaxBackup int    `yaml:"max_backup"`
	MaxSize   int    `yaml:"max_size"`
	MaxAge    int    `yaml:"max_age"`
	Compress  bool   `yaml:"compress"`
}
