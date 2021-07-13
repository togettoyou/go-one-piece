package errno

// 例: 10001
// 第一位的1表示服务级错误 (1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的)
// 第二位至第三位的00表示服务模块代码
// 最后两位01表示具体错误代码

var (
	OK                  = &Errno{Code: 0, Message: "成功"}
	InternalServerError = &Errno{Code: 10001, Message: "服务器异常"}

	ErrValidation = &Errno{Code: 20001, Message: "参数校验失败"}
	ErrBind       = &Errno{Code: 20002, Message: "参数绑定异常"}
	ErrUnknown    = &Errno{Code: 20003, Message: "未知错误"}

	ErrNotLogin     = &Errno{Code: 20101, Message: "请登录"}
	ErrTokenExpired = &Errno{Code: 20102, Message: "令牌已过期"}
	ErrTokenInvalid = &Errno{Code: 20103, Message: "令牌无效"}
	ErrTokenFailure = &Errno{Code: 20104, Message: "令牌验证失败"}
)
