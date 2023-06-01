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
	Success      OpsError = NewOpsError(200, "success")
	ServerErr    OpsError = NewOpsError(500, "server error")
	NoPermission OpsError = NewOpsError(501, "无权限")
	NotLogin     OpsError = NewOpsError(401, "未登录")

	TokenExpire  OpsError = NewOpsError(4001, "token过期，请重新登录")
	TokenInvalid OpsError = NewOpsError(4002, "token无效")

	UserDeny    OpsError = NewOpsError(1001, "用户已被锁定，请联系管理员")
	PasswdWrong OpsError = NewOpsError(1002, "登录失败，请输入正确的账号和密码")
	CaptchaErr  OpsError = NewOpsError(1003, "验证码错误")
	RegisterErr OpsError = NewOpsError(1004, "注册失败")
	UserExist   OpsError = NewOpsError(1005, "用户已存在")

	GetResourceErr OpsError = NewOpsError(2001, "获取资源失败")
	OperatorErr    OpsError = NewOpsError(2001, "操作失败")

	ParamErr OpsError = NewOpsError(3001, "参数错误")
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
