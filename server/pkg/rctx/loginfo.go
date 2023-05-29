package rctx

type LogInfo struct {
	LogModule string // 属于那个模块的日志
	LogRes    bool   //是否记录请求返回结果
}

func NewLogInfo() *LogInfo {
	return &LogInfo{LogModule: "系统默认", LogRes: false}
}

func (l *LogInfo) WithModule(module string) *LogInfo {
	l.LogModule = module
	return l
}

// WithLogRes 是否记录请求返回结果，默认false
func (l *LogInfo) WithLogRes(flag bool) *LogInfo {
	l.LogRes = flag
	return l
}
