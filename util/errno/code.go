package errno

// 例: 10001
// 第一位的1表示服务级错误 (1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的)
// 第二位至第三位的00表示服务模块代码
// 最后两位01表示具体错误代码

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "成功"}
	InternalServerError = &Errno{Code: 10001, Message: "服务器异常"}
	ErrBind             = &Errno{Code: 10002, Message: "参数绑定异常"}
	ErrSQLUnknown       = &Errno{Code: 10003, Message: "sql执行出错"}
	ErrUnknown          = &Errno{Code: 10004, Message: "未知错误"}

	ErrValidation   = &Errno{Code: 20001, Message: "参数校验失败"}
	ErrNotLogin     = &Errno{Code: 20002, Message: "请登录"}
	ErrTokenExpired = &Errno{Code: 20003, Message: "令牌已过期"}
	ErrTokenInvalid = &Errno{Code: 20004, Message: "令牌无效"}
	ErrTokenFailure = &Errno{Code: 20005, Message: "令牌验证失败"}
	ErrPermission   = &Errno{Code: 20006, Message: "权限不足"}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20101, Message: "用户不存在"}
	ErrUserExisting      = &Errno{Code: 20102, Message: "用户已存在"}
	ErrPasswordIncorrect = &Errno{Code: 20103, Message: "密码不正确"}

	// role errors
	ErrRoleNotFound     = &Errno{Code: 20201, Message: "角色不存在"}
	ErrRoleExisting     = &Errno{Code: 20202, Message: "角色已存在"}
	ErrUserRoleNotFound = &Errno{Code: 20203, Message: "用户未分配角色"}

	// casbin errors
	ErrUpdateCasbinRoleApi = &Errno{Code: 20301, Message: "更新角色权限失败"}
	ErrSetCasbinUserRole   = &Errno{Code: 20302, Message: "设置用户角色失败"}
)
