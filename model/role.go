package model

import (
	"go-one-server/util/errno"
	"gorm.io/gorm"
)

// Role 角色
type Role struct {
	Model
	RoleInfo
}

type RoleInfo struct {
	RoleKey  string `json:"role_key" gorm:"unique;not null;size:32;comment:角色代码" binding:"required,alphanum,min=4,max=32" example:"角色代码"`
	RoleName string `json:"role_name" gorm:"unique;not null;type:varchar(20);comment:角色名字" binding:"required,min=2,max=20" example:"角色名"`
	Remark   string `json:"remark" gorm:"comment:备注" example:"备注"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	var count int64
	tx.Model(&Role{}).Where("role_key = ? OR role_name = ?", r.RoleKey, r.RoleName).Count(&count)
	if count > 0 {
		return errno.ErrRoleExisting
	}
	return nil
}

func (r *Role) Create() error {
	return db.Create(r).Error
}

func GetRoleList(page, pageSize int) (data *PaginationQ, err error) {
	var roles []*Role
	var total int64
	err = db.Model(&Role{}).Scopes(Count(&total)).Scopes(Paginate(&page, &pageSize)).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return &PaginationQ{
		PageSize: pageSize,
		Page:     page,
		Data:     roles,
		Total:    total,
	}, nil
}

func FindRoleByKey(roleKey string) (*Role, error) {
	var role Role
	if err := db.Where(map[string]interface{}{"role_key": roleKey}).
		Take(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ErrRoleNotFound
		}
		return nil, err
	}
	return &role, nil
}

func DelRole(roleKey string) error {
	return db.Where("role_key = ?", roleKey).
		Delete(&Role{}).Error
}
