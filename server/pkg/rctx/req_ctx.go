package rctx

import (
	"fmt"
	"github.com/lbemi/lbemi/apps/system/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/emicklei/go-restful/v3"

	"github.com/lbemi/lbemi/pkg/restfulx"
)

type HandlerFunc func(ctx *ReqCtx)

type ReqCtx struct {
	Keys     map[string]any
	mu       sync.RWMutex
	Request  *restful.Request
	Response *restful.Response

	RequirePermission *Permission

	LoginAccount *entity.User
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
		NoRes:             false,
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

// WithNoRes websocket,文件下载等无需返回结果
func (rc *ReqCtx) WithNoRes() *ReqCtx {
	rc.NoRes = true
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
		err := ApplyHandlerInterceptor(afterHandlers, rc)
		if err != nil {
			return
		}
	}()
	// 如果rc.Response 为nil，则panic
	restfulx.ErrNotTrue(rc.Request != nil, restfulx.NewErr("Request == nil"))
	rc.ReqParam = 0
	rc.ResData = nil
	err := ApplyHandlerInterceptor(beforeHandlers, rc)
	if err != nil {
		return
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
func (rc *ReqCtx) ClientIP() string {
	r := rc.Request
	ip := r.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = r.Request.Header.Get("X-real-ip")
	}

	if ip == "" {
		return "127.0.0.1"
	}

	return ip
}

// GetLocation identifies the location of a user based on their IP address.

func (rc *ReqCtx) ShouldBind(data interface{}) {
	if err := rc.Request.ReadEntity(data); err != nil {
		restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	}
}

// GetPageQueryParam 获取分页参数
func (rc *ReqCtx) GetPageQueryParam() *entity2.PageParam {
	return &entity2.PageParam{Page: rc.QueryDefaultInt("page", 0), Limit: rc.QueryDefaultInt("limit", 0)}
}

// QueryDefaultInt 获取查询参数中指定参数值，并转为int
func (rc *ReqCtx) QueryDefaultInt(key string, defaultInt int) int {
	qv := rc.Request.QueryParameter(key)
	if qv == "" {
		return defaultInt
	}
	qvi, err := strconv.Atoi(qv)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return qvi
}

func (rc *ReqCtx) QueryDefault(key string, defaultStr string) string {
	qv := rc.Request.QueryParameter(key)
	if qv == "" {
		return defaultStr
	}
	return qv
}

// Query get query param
func (rc *ReqCtx) Query(key string) string {
	return rc.Request.QueryParameter(key)
}

// QueryCloud get cloud name ,if cloud is empty return param error
func (rc *ReqCtx) QueryCloud() string {
	cloud := rc.Request.QueryParameter("cloud")
	if cloud == "" {
		restfulx.ErrIsNil(fmt.Errorf("cloud is empty"), restfulx.ClusterNotSet)
	}
	return cloud
}

// QueryParamUint8 Query
func (rc *ReqCtx) QueryParamUint8(key string) uint8 {
	str := rc.Request.QueryParameter(key)
	if str == "" {
		return uint8(0)
	}
	i, err := strconv.Atoi(str)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return uint8(i)
}

// QueryParamInt8 Query
func (rc *ReqCtx) QueryParamInt8(key string) int8 {
	str := rc.Request.QueryParameter(key)
	if str == "" {
		return int8(0)
	}
	i, err := strconv.Atoi(str)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return int8(i)
}

// PathParamInt 获取路径参数
func (rc *ReqCtx) PathParamInt(key string) int {
	value, err := strconv.Atoi(rc.Request.PathParameter(key))
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return value
}

// PathParamInt64 获取路径参数
func (rc *ReqCtx) PathParamInt64(key string) int64 {
	value, err := strconv.ParseInt(rc.Request.PathParameter(key), 10, 64)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return value
}

// PathParamUint64 获取路径参数
func (rc *ReqCtx) PathParamUint64(key string) uint64 {
	value, err := strconv.ParseUint(rc.Request.PathParameter(key), 10, 64)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return value
}
func (rc *ReqCtx) PathParam(pm string) string {
	return rc.Request.PathParameter(pm)
}

func (rc *ReqCtx) FormFile(key string) []byte {
	_, fileHeader, err := rc.Request.Request.FormFile(key)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	file, err := fileHeader.Open()
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	bytes, err := io.ReadAll(file)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return bytes
}

func (rc *ReqCtx) PostForm(key string) string {
	return rc.Request.Request.PostFormValue(key)
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
