package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"net/http"
)

func SuccessRes(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessX(data))
}

func ErrorRes(c *gin.Context, err interface{}) {
	switch t := err.(type) {
	case GinError:
		c.JSON(http.StatusOK, Error(t))
	case error:
		c.JSON(http.StatusOK, ServerError())
		log.Logger.Error(err)
	case string:
		c.JSON(http.StatusOK, ServerError())
		log.Logger.Error(err)
	default:
		log.Logger.Error(err)
	}
}
