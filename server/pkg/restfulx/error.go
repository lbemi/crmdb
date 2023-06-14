package restfulx

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"runtime/debug"
)

var (
	Success      *OpsError = NewOpsErrCode(200, "success")
	ServerErr    *OpsError = NewOpsErrCode(500, "服务器内部错误")
	NoPermission *OpsError = NewOpsErrCode(403, "无权限")
	NotLogin     *OpsError = NewOpsErrCode(401, "未登录")

	TokenExpire  *OpsError = NewOpsErrCode(4001, "token过期，请重新登录")
	TokenInvalid *OpsError = NewOpsErrCode(4002, "token无效")

	UserDeny      *OpsError = NewOpsErrCode(1001, "用户已被锁定，请联系管理员")
	PasswdWrong   *OpsError = NewOpsErrCode(1002, "登录失败，请输入正确的账号和密码")
	CaptchaErr    *OpsError = NewOpsErrCode(1003, "验证码错误")
	RegisterErr   *OpsError = NewOpsErrCode(1004, "注册失败")
	UserExist     *OpsError = NewOpsErrCode(1005, "用户已存在")
	ResourceExist *OpsError = NewOpsErrCode(1006, "资源已存在")

	GetResourceErr *OpsError = NewOpsErrCode(2001, "获取资源失败")
	OperatorErr    *OpsError = NewOpsErrCode(2001, "操作失败")

	ParamErr *OpsError = NewOpsErrCode(3001, "参数错误")

	//ks8相关
	RegisterClusterErr *OpsError = NewOpsErrCode(5001, "导入集群失败，请检查配置文件")
	ClusterUnHealth    *OpsError = NewOpsErrCode(5002, "集群异常，请检查")
)

type OpsError struct {
	code    int16
	message string
}

func (oe *OpsError) Code() int16 {
	return oe.code
}

func (oe *OpsError) Error() string {
	return oe.message
}

func NewErr(msg string) *OpsError {
	return &OpsError{code: ServerErr.code, message: msg}
}
func NewOpsErrCode(code int16, err string) *OpsError {
	return &OpsError{code: code, message: err}
}

func Error(oe *OpsError) *Response {
	return &Response{Code: oe.Code(), Message: oe.Error()}
}

func ServerError() *Response {
	return Error(ServerErr)
}

func ErrNotNil(err error, msg string, param ...any) {
	if err != nil {
		//log.Logger.Error(message)
		panic(NewErr(fmt.Sprintf(msg, param...)))
	}
}

func ErrIsNil(err error, oe *OpsError) {
	if err != nil {
		panic(oe)
	}
}

func ErrNotNilDebug(err error, oe *OpsError) {
	if err != nil {
		log.Logger.Error(err)
		debug.PrintStack()
		panic(oe)
	}
}

func ErrNotTrue(exp bool, err *OpsError) {
	if !exp {
		panic(err)
	}
}
