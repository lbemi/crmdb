package log

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/model/configs"
	"os"
	"path/filepath"
)

type Configuration struct {
	LogType  string
	LogFile  string
	LogLevel string
	IsFile   bool

	RotateMaxSize    int
	RotateMaxAge     int
	RotateMaxBackups int
	Compress         bool
}

type LoggerInterface interface {
	Info(args ...interface{})
	Infof(f string, args ...interface{})
	Error(args ...interface{})
	Errorf(f string, args ...interface{})
	Warn(args ...interface{})
	Warnf(f string, args ...interface{})
}

var (
	Logger LoggerInterface
)

func Register(config *configs.Log) {
	//logType, logDir, logLevel string
	// 支持 INFO, WARN 和 ERROR，默认为 INFO
	if config.IsFile {
		createRootDir(config)
	}

	Logger, _ = NewZapLogger(Configuration{
		IsFile:           config.IsFile,
		LogType:          config.Format,
		LogFile:          filepath.Join(config.RootDir, config.FileName), // 使用文件类型时生效
		LogLevel:         config.Level,                                   // access 的 log 只会有 info
		RotateMaxSize:    config.MaxSize,
		RotateMaxAge:     config.MaxAge,
		RotateMaxBackups: config.MaxBackup,
	})
}

func createRootDir(config *configs.Log) {
	if ok, _ := pathExists(config.RootDir); !ok {
		err := os.Mkdir(config.RootDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
