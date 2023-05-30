package rctx

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"strconv"
)

func QueryPageParam(c *gin.Context) *model.PageParam {
	return &model.PageParam{Page: QueryInt(c, "page", 1), Limit: QueryInt(c, "limit", 1)}
}

func QueryInt(c *gin.Context, key string, defaultValue int) int {
	res := c.Query(key)
	if res == "" {
		return defaultValue
	}
	intRes, err := strconv.Atoi(res)
	restfulx.ErrIsNil(err, "get page param error")
	return intRes
}

func ShouldBind(rc *ReqCtx, data any) {
	if err := rc.Request.ReadEntity(data); err != nil {
		panic(restfulx.NewErr(err.Error()))
	}
}
