package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/model"
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
	ErrIsNil(err, "get page param error")
	return intRes
}
