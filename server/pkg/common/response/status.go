package response

// 自定义错误
const (
	StatusOK                  = 200
	StatusInternalServerError = 500

	ErrCodeSuccess   = 2005
	ErrOperateFailed = 2006 // 操作失败

	ErrCodeUserExist  = 2001 //数据库错误码
	ErrCodeServerBusy = 2002
	ErrCodeNotFount   = 2003
	ErrCodeFount      = 2004

	ErrCodeParameter         = 1001
	ErrCodeNotLogin          = 1002 //用户未登录
	ErrCodeUserForbidden     = 1003 //用户已被禁用
	ErrCodeUserNotExist      = 1004
	ErrCodeUserOrPasswdWrong = 1005
	ErrCodeRegisterFail      = 1006
	ErrCodeGenCaptchaFail    = 1007
	ErrCaptcha               = 100 // 验证码错误

	ErrorTokenExpired = 9001
	InvalidToken      = 9002

	LoginSuccess = 2000
	AddSuccess   = 3000

	ClusterNoHealth = 4000

	NoPermission = 5000
)

// GetMessage 根据错误码返回错误信息
func GetMessage(code int) (message string) {
	switch code {
	case ErrOperateFailed:
		message = "操作失败"
	case ErrCodeGenCaptchaFail:
		message = "验证码生成失败"
	case ErrCaptcha:
		message = "验证码错误"
	case ErrCodeRegisterFail:
		message = "注册失败"
	case ErrCodeNotFount:
		message = "数据不存在"
	case ErrCodeFount:
		message = "数据已存在"
	case StatusInternalServerError:
		message = "服务器内部错误"
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeUserForbidden:
		message = "用户已被禁用,请联系管理员"
	case NoPermission:
		message = "no permission"
	case StatusOK:
		message = "操作成功"
	case AddSuccess:
		message = "添加成功"
	case InvalidToken:
		message = "无效的token"
	case LoginSuccess:
		message = "登录成功"
	case ErrCodeParameter:
		message = "参数错误"
	case ErrCodeUserExist:
		message = "用户名已存在"
	case ErrCodeServerBusy:
		message = "服务器繁忙"
	case ErrCodeUserNotExist:
		message = "用户不存在"
	case ErrCodeUserOrPasswdWrong:
		message = "账号或密码错误"
	case ErrCodeNotLogin:
		message = "用户未登录"
	case ErrorTokenExpired:
		message = "token 失效"
	case ClusterNoHealth:
		message = "集群异常"
	default:
		message = "未知错误"
	}
	return
}
