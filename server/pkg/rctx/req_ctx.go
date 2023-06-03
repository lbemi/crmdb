package rctx

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"sync"
	"time"
)

type HandlerFunc func(ctx *ReqCtx)

type ReqCtx struct {
	Keys     map[string]any
	mu       sync.RWMutex
	Request  *restful.Request
	Response *restful.Response

	RequirePermission *Permission

	LoginAccount *sys.User
	LogInfo      *LogInfo

	ReqParam any
	ResData  any
	Err      any

	Timed int64
	NoRes bool

	Handler HandlerFunc
}

func NewReqCtx(request *restful.Request, response *restful.Response) *ReqCtx {
	return &ReqCtx{
		Request:           request,
		Response:          response,
		LogInfo:           NewLogInfo(),
		RequirePermission: NewPermission(),
	}
}

// WithLog 指定log日志属于那个模块
func (rc *ReqCtx) WithLog(module string) *ReqCtx {
	rc.LogInfo.WithModule(module)
	return rc
}

// WithToken 是否启用校验token，默认启用
func (rc *ReqCtx) WithToken(flag bool) *ReqCtx {
	rc.RequirePermission.WithToken(flag)
	return rc
}

// WithCasbin 是否启用casbin鉴权，默认启用
func (rc *ReqCtx) WithCasbin(flag bool) *ReqCtx {
	rc.RequirePermission.WithCasbin(flag)
	return rc
}

func (rc *ReqCtx) WithHandle(handler HandlerFunc) *ReqCtx {
	rc.Handler = handler
	return rc
}

// Do 处理handles
func (rc *ReqCtx) Do() {
	defer func() {
		if err := recover(); err != nil {
			rc.Err = err
			restfulx.ErrorRes(rc.Response, err)
		}
		ApplyHandlerInterceptor(afterHandlers, rc)
	}()
	// 如果rc.Response 为nil，则panic
	restfulx.ErrNotTrue(rc.Request != nil, restfulx.NewErr("Request == nil"))
	//util.IsTrue(rc.Response != nil, "Response == nil")
	rc.ReqParam = 0
	rc.ResData = nil
	err := ApplyHandlerInterceptor(beforeHandlers, rc)
	if err != nil {
		panic(err)
	}
	begin := time.Now()
	rc.Handler(rc)
	rc.Timed = time.Since(begin).Milliseconds()
	if !rc.NoRes {
		restfulx.SuccessRes(rc.Response, rc.ResData)
	}
}

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (rc *ReqCtx) Set(key string, value any) {
	rc.mu.Lock()
	if rc.Keys == nil {
		rc.Keys = make(map[string]any)
	}

	rc.Keys[key] = value
	rc.mu.Unlock()
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (rc *ReqCtx) Get(key string) (value any, exists bool) {
	rc.mu.RLock()
	value, exists = rc.Keys[key]
	rc.mu.RUnlock()
	return
}

type HandlerInterceptorFunc func(ctx *ReqCtx) error
type HandlerInterceptors []HandlerInterceptorFunc

var (
	beforeHandlers HandlerInterceptors
	afterHandlers  HandlerInterceptors
)

// UseBeforeHandlerInterceptor 添加前置处理函数
func UseBeforeHandlerInterceptor(h HandlerInterceptorFunc) {
	beforeHandlers = append(beforeHandlers, h)
}

// UserAfterHandlerInterceptor 添加后置处理函数
func UserAfterHandlerInterceptor(h HandlerInterceptorFunc) {
	afterHandlers = append(afterHandlers, h)
}

// ApplyHandlerInterceptor 执行拦截器
func ApplyHandlerInterceptor(handlers HandlerInterceptors, rc *ReqCtx) error {
	for _, handler := range handlers {
		if err := handler(rc); err != nil {
			return err
		}
	}
	return nil
}
