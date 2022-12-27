package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"reflect"
)

type Result struct {
	data  []interface{} `json:"data"`
	total int           `json:"total"`
}

func PageQuery(c *gin.Context, data interface{}, page, limit int) {
	var pageQuery Result
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice {
		//res.data = data.([]interface{})[(page-1)*limit : page*limit]
		//res.total = len(res.data)
		//response.Success(c, response.StatusOK, res)
		data2 := data.([]any)
		// 处理分页
		pageQuery.total = len(data2)

		if pageQuery.total <= limit {
			pageQuery.data = data2
		} else if page*limit >= pageQuery.total {
			pageQuery.data = data2[(page-1)*limit : pageQuery.total]
		} else {
			pageQuery.data = data2[(page-1)*limit : page*limit]
		}
		response.Success(c, response.StatusOK, pageQuery)
	} else {
		response.Fail(c, response.StatusInternalServerError)
	}
}
