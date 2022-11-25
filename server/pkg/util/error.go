package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
)

func WithErrorLog(err error) {
	if err != nil {
		log.Logger.Error(err)
		return
	}
}

func GinError(c *gin.Context, err error, code int) {
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, code)
		panic(err)
	}
}