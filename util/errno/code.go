package errno

// 例: 10001
// 第一位的1表示服务级错误 (1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的)
// 第二位至第三位的00表示服务模块代码
// 最后两位01表示具体错误代码

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrUnknown          = &Errno{Code: 10003, Message: "Unknown Error"}

	ErrValidation   = &Errno{Code: 20001, Message: "Validation failed."}
	ErrNotLogin     = &Errno{Code: 20002, Message: "Please login."}
	ErrTokenExpired = &Errno{Code: 20003, Message: "The token was expired."}
	ErrTokenInvalid = &Errno{Code: 20004, Message: "The token was invalid."}
	ErrTokenFailure = &Errno{Code: 20005, Message: "The token validation failure."}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20101, Message: "The user was not found."}
	ErrUserExisting      = &Errno{Code: 20102, Message: "The user already exists."}
	ErrPasswordIncorrect = &Errno{Code: 20103, Message: "The password was incorrect."}
)
