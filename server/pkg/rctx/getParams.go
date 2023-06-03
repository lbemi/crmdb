package rctx

import (
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"strconv"
)

// GetPageQueryParam 获取分页参数
func GetPageQueryParam(rc *ReqCtx) *model.PageParam {
	return &model.PageParam{Page: QueryDefaultInt(rc, "page", 1), Limit: QueryDefaultInt(rc, "limit", 10)}
}

// QueryDefaultInt 获取查询参数中指定参数值，并转为int
func QueryDefaultInt(rc *ReqCtx, key string, defaultInt int) int {
	qv := rc.Request.QueryParameter(key)
	if qv == "" {
		return defaultInt
	}
	qvi, err := strconv.Atoi(qv)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return qvi
}

func QueryDefault(rc *ReqCtx, key string, defaultStr string) string {
	qv := rc.Request.QueryParameter(key)
	if qv == "" {
		return defaultStr
	}
	return qv
}

// QueryParam QueryParam
func QueryParam(rc *ReqCtx, key string) string {
	return rc.Request.QueryParameter(key)
}

// PathParamInt 获取路径参数
func PathParamInt(rc *ReqCtx, key string) int {
	value, err := strconv.Atoi(rc.Request.PathParameter(key))
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return value
}

// PathParamUint64 获取路径参数
func PathParamUint64(rc *ReqCtx, key string) uint64 {
	value, err := strconv.ParseUint(rc.Request.PathParameter(key), 10, 64)
	restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	return value
}
func PathParam(rc *ReqCtx, pm string) string {
	return rc.Request.PathParameter(pm)
}

func ShouldBind(rc *ReqCtx, data any) {
	if err := rc.Request.ReadEntity(data); err != nil {
		restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
	}
}
