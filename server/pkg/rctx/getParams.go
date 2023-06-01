package rctx

import (
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"strconv"
)

// GetPageQueryParam 获取分页参数
func GetPageQueryParam(rc *ReqCtx) *model.PageParam {
	return &model.PageParam{Page: QueryInt(rc, "page", 1), Limit: QueryInt(rc, "limit", 10)}
}

// QueryInt 获取查询参数中指定参数值，并转为int
func QueryInt(rc *ReqCtx, key string, defaultInt int) int {
	qv := rc.Request.QueryParameter(key)
	if qv == "" {
		return defaultInt
	}
	qvi, err := strconv.Atoi(qv)
	restfulx.ErrIsNilRes(err, restfulx.ParamErr)
	return qvi
}

// QueryParam QueryParam
func QueryParam(rc *ReqCtx, key string) string {
	return rc.Request.QueryParameter(key)
}

// ParamInt 获取路径参数
func ParamInt(rc *ReqCtx, key string) int {
	value, err := strconv.Atoi(rc.Request.PathParameter(key))
	restfulx.ErrIsNilRes(err, restfulx.ParamErr)
	return value
}

// ParamUint64 获取路径参数
func ParamUint64(rc *ReqCtx, key string) uint64 {
	value, err := strconv.ParseUint(rc.Request.PathParameter(key), 10, 64)
	restfulx.ErrIsNilRes(err, restfulx.ParamErr)
	return value
}
func PathParam(rc *ReqCtx, pm string) string {
	return rc.Request.PathParameter(pm)
}

func ShouldBind(rc *ReqCtx, data any) {
	if err := rc.Request.ReadEntity(data); err != nil {
		restfulx.ErrIsNilRes(err, restfulx.ParamErr)
	}
}
