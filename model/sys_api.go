package model

import (
	"go-one-server/util/tools"
	"gorm.io/gorm"
)

type ApiMethod string

const (
	POST   ApiMethod = "POST"
	GET              = "GET"
	PATCH            = "PATCH"
	PUT              = "PUT"
	DELETE           = "DELETE"
)

// Api 行为权限表
type Api struct {
	CreatedAt   tools.FormatTime `json:"created_at"`
	UpdatedAt   tools.FormatTime `json:"-"`
	DeletedAt   gorm.DeletedAt   `json:"-" gorm:"index"`
	Path        string           `json:"path" gorm:"primarykey;comment:api路径"`
	Method      ApiMethod        `json:"method" gorm:"primarykey;type:varchar(6);default:POST;comment:请求方法"`
	Description string           `json:"description" gorm:"comment:api中文描述"`
}
