package bootstrap

import (
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	level   zapcore.Level
	options []zap.Option
)

func InitializeLog() *zap.Logger {
	//createRoot
}

func createRootDir() {
	if ok, _ := util.PathExists(global.App.Config.RootDir); !ok {
		_ = os.Mkdir(global.App.Config.RootDir, os.ModePerm)
	}
}
func SetLogLevel() {
	switch global.App.Config.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	}
}
