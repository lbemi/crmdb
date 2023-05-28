package rctx

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/ginx"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/util"
	"net/http"
	"time"
)

type HandlerFunc func(ctx *ReqCtx)

type ReqCtx struct {
	GinCtx            *gin.Context
	RequirePermission *Permission

	LoginAccount *sys.User
	LogInfo      *LogInfo

	ReqParam any
	ResData  any
	Err      any

	Timed int64
	NoRes bool
}

func NewReqCtx(ginCtx *gin.Context) *ReqCtx {
	return &ReqCtx{GinCtx: ginCtx, LogInfo: NewLogInfo(), RequirePermission: NewPermission()}
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

// Handle 处理handle
func (rc *ReqCtx) Handle(handler HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rc.Err = err
			rc.GinCtx.JSON(http.StatusOK, err)
		}
		ApplyHandlerInterceptor(afterHandlers, rc)
	}()
	// 如果rc.GinCtx 为nil，则panic
	util.IsTrue(rc.GinCtx != nil, "ginContext == nil")

	rc.ReqParam = 0
	rc.ResData = nil
	err := ApplyHandlerInterceptor(beforeHandlers, rc)
	if err != nil {
		panic(err)
	}
	begin := time.Now()
	handler(rc)
	rc.Timed = time.Since(begin).Milliseconds()
	if !rc.NoRes {
		ginx.SuccessRes(rc.GinCtx, rc.ResData)
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
