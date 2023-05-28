package ginx

import (
	"encoding/json"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"net/http"
)

type GinError struct {
	code int16
	err  string
}

func NewGinErrorCode(code int16, err string) GinError {
	return GinError{code: code, err: err}
}

var (
	Success       GinError = NewGinErrorCode(200, "success")
	ServerErr     GinError = NewGinErrorCode(500, "server error")
	PermissionErr GinError = NewGinErrorCode(501, "no permission")
	GinErr        GinError = NewGinErrorCode(400, "gin error")
)

func (ge GinError) Code() int16 {
	return ge.code
}

func (ge GinError) Error() string {
	return ge.err
}

func NewGinErr(msg string) GinError {
	return GinError{code: GinErr.code, err: msg}
}

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
	return &Response{Code: Success.code, Message: SuccessMsg, Data: data}
}

func Error(ge GinError) *Response {
	return &Response{Code: ge.Code(), Message: ge.Error()}
}

func ServerError() *Response {
	return Error(ServerErr)
}

func PermissionError() *Response {
	return Error(PermissionErr)
}

func ErrIsNil(err error, msg string, param ...any) {
	if err != nil {
		log.Logger.Error(err)
		panic(NewGinErr(fmt.Sprintf(msg, param...)))
	}
}
