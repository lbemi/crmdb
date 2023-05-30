package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"net/http"
)

const SuccessMsg = "success"

type Response struct {
	Code    int16       `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func (r *Response) ToJson() string {
	marshalData, err := json.Marshal(r.Data)
	if err != nil {
		fmt.Println("marshal data to json failed")
	}
	return string(marshalData)
}
func (r *Response) IsSuccess() bool {
	return r.Code == http.StatusOK
}

func SuccessX(data interface{}) *Response {
	return &Response{Code: StatusOK, Message: SuccessMsg, Data: data}
}

func Error(ge restfulx.OpsError) *Response {
	return &Response{Code: ge.Code(), Message: ge.Error()}
}

func Success(c *gin.Context, code int16, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		GetMessage(code),
	})
}

func Fail(c *gin.Context, code int16) {
	c.JSON(http.StatusOK, Response{
		code,
		nil,
		GetMessage(code),
	})
}

func FailWithMessage(c *gin.Context, code int16, message string) {
	c.JSON(http.StatusOK, Response{
		code,
		nil,
		message,
	})
}
