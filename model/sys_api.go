package model

type ApiMethod string

const (
	POST   ApiMethod = "POST"
	GET              = "GET"
	PATCH            = "PATCH"
	PUT              = "PUT"
	DELETE           = "DELETE"
)

// Action 行为权限表
type Api struct {
	Model
	Path        string    `json:"path" gorm:"comment:api路径"`
	Method      ApiMethod `json:"method" gorm:"type:varchar(6);default:POST;comment:请求方法"`
	Description string    `json:"description" gorm:"comment:api中文描述"`
}
