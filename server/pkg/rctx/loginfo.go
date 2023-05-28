package rctx

type LogInfo struct {
	LogModule string // 属于那个模块的日志
}

func NewLogInfo() *LogInfo {
	return &LogInfo{LogModule: "系统默认"}
}

func (l *LogInfo) WithModule(module string) *LogInfo {
	l.LogModule = module
	return l
}
