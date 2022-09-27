package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		GetMessage(code),
	})
}
func Fail(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		code,
		nil,
		GetMessage(code),
	})
}

func FailWithMessage(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		code,
		nil,
		message,
	})
}
