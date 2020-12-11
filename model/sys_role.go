package model

import (
	"go-one-server/util/tools"
	"gorm.io/gorm"
)

// Role 角色
type Role struct {
	CreatedAt tools.FormatTime `json:"created_at"`
	UpdatedAt tools.FormatTime `json:"-"`
	DeletedAt gorm.DeletedAt   `json:"-" gorm:"index"`
	RoleID    string           `json:"role_id" gorm:"unique;not null;primarykey;type:varchar(32);comment:角色ID"`
	RoleName  string           `json:"role_name" gorm:"unique;not null;type:varchar(20);comment:角色名字"`
}
