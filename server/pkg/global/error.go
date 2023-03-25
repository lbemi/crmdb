package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError  CustomError
	ValidatorError CustomError
}

var Error = CustomErrors{
	BusinessError:  CustomError{4000, "业务错误"},
	ValidatorError: CustomError{4200, "请求参数错误"},
}
