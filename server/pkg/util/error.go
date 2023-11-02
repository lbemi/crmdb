package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/global"

	"github.com/lbemi/lbemi/pkg/common/response"
)

func WithErrorLog(err error) {
	if err != nil {
		global.Logger.Error(err)
		return
	}
}

func GinError(c *gin.Context, err error, code int16) {
	if err != nil {
		response.Fail(c, code)
		panic(err)
	}
}

func HandleError(err error) {
	if err != nil {
		global.Logger.Error(err)
		panic(err)
	}
}
