package restfulx

import (
	"encoding/json"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"net/http"
)

type OpsError struct {
	code    int16
	message string
}

func NewOpsError(code int16, err string) OpsError {
	return OpsError{code: code, message: err}
}

var (
	Success       OpsError = NewOpsError(200, "success")
	ServerErr     OpsError = NewOpsError(500, "server error")
	PermissionErr OpsError = NewOpsError(501, "no permission")
	GinErr        OpsError = NewOpsError(400, "gin error")
	NotLogin      OpsError = NewOpsError(401, "please login")

	TokenExpire  OpsError = NewOpsError(4001, "token expired")
	TokenInvalid OpsError = NewOpsError(4002, "token invalid")

	UserDeny    OpsError = NewOpsError(1001, "The user has been disabled. Please contact the administrator")
	PasswdWrong OpsError = NewOpsError(1002, "login failed. please input right password or user")
	CaptchaErr  OpsError = NewOpsError(1003, "captcha error")
	RegisterErr OpsError = NewOpsError(1004, "register error")
	UserExist   OpsError = NewOpsError(1005, "user existed")
)

func (oe OpsError) Code() int16 {
	return oe.code
}

func (oe OpsError) Error() string {
	return oe.message
}

func NewErr(msg string) OpsError {
	return OpsError{code: ServerErr.code, message: msg}
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

func Error(oe OpsError) *Response {
	return &Response{Code: oe.Code(), Message: oe.Error()}
}

func ServerError() *Response {
	return Error(ServerErr)
}

func PermissionError() *Response {
	return Error(PermissionErr)
}

func ErrIsNil(err error, msg string, param ...any) {
	if err != nil {
		//log.Logger.Error(message)
		panic(NewErr(fmt.Sprintf(msg, param...)))
	}
}

func ErrIsNilRes(err error, oe OpsError) {
	if err != nil {
		log.Logger.Error(err)
		panic(oe)
	}
}

func ErrNotTrue(exp bool, err OpsError) {
	if !exp {
		//log.Logger.Error(message.message)
		panic(err)
	}
}
