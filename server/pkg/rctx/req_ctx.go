package rctx

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/pkg/ginx"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/util"
	"time"
)

type HandlerFunc func(ctx *ReqCtx)

type ReqCtx struct {
	Keys     map[string]any
	Request  *restful.Request
	Response *restful.Response

	RequirePermission *Permission

	LoginAccount *sys.User
	LogInfo      *LogInfo

	ReqParam any
	ResData  any
	Err      any

	Timed time.Duration
	NoRes bool

	Handler HandlerFunc
}

func NewReqCtx() *ReqCtx {
	return &ReqCtx{
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

// Handle 处理handle
func (rc *ReqCtx) Do() restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		rc.Request = request
		rc.Response = response
		defer func() {
			if err := recover(); err != nil {
				rc.Err = err
				ginx.ErrorRes(rc.Response, err)
			}
			ApplyHandlerInterceptor(afterHandlers, rc)
		}()
		// 如果rc.GinCtx 为nil，则panic
		util.IsTrue(rc.Response != nil, "Response == nil")

		rc.ReqParam = 0
		rc.ResData = nil
		err := ApplyHandlerInterceptor(beforeHandlers, rc)
		if err != nil {
			panic(err)
		}
		begin := time.Now()
		rc.Handler(rc)
		rc.Timed = time.Since(begin)
		if !rc.NoRes {
			ginx.SuccessRes(rc.Response, rc.ResData)
		}
	}
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
