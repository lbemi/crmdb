package bootstrap

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level   zapcore.Level
	options []zap.Option
	allCore []zapcore.Core
)

func InitializeLog() *zap.Logger {
	if global.App.Config.IsFile {
		createRootDir()
	}
	SetLogLevel()
	if global.App.Config.Log.ShowLine {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(0))
	}
	core := getZapCore()
	return zap.New(core).WithOptions(options...)
}

func createRootDir() {
	if ok, _ := util.PathExists(global.App.Config.RootDir); !ok {
		err := os.Mkdir(global.App.Config.RootDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
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
func getLogWriter() zapcore.WriteSyncer {

	file := &lumberjack.Logger{
		Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.FileName,
		MaxAge:     global.App.Config.MaxAge,
		MaxSize:    global.App.Config.MaxSize,
		MaxBackups: global.App.Config.MaxBackup,
		Compress:   global.App.Config.Compress,
	}
	return zapcore.AddSync(file)
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.Env + "." + l.String() + ":")
	}
	if global.App.Config.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	consoleDebugging := zapcore.Lock(os.Stdout)
	console := zapcore.NewCore(encoder, consoleDebugging, level)
	allCore = append(allCore, console)
	if global.App.Config.IsFile {
		file := zapcore.NewCore(encoder, getLogWriter(), level)
		allCore = append(allCore, file)
	}
	//if global.App.Config.Database.EnableFileLogWrite {
	//	gorm := zapcore.NewCore(encoder, getGormLogWriter(), level)
	//	allCore = append(allCore, gorm)
	//}
	core := zapcore.NewTee(allCore...)
	return core
}
