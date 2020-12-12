package model

import (
	"go-one-server/util/errno"
	"go-one-server/util/tools"
	"gorm.io/gorm"
)

// Role 角色
type Role struct {
	CreatedAt tools.FormatTime `json:"created_at"`
	UpdatedAt tools.FormatTime `json:"-"`
	DeletedAt gorm.DeletedAt   `json:"-" gorm:"index"`
	RoleInfo
}

type RoleInfo struct {
	RoleID   string `json:"role_id" gorm:"unique;not null;primarykey;type:varchar(32);comment:角色ID"`
	RoleName string `json:"role_name" gorm:"unique;not null;type:varchar(20);comment:角色名字"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	var count int64
	tx.Model(&Role{}).Where("role_name = ?", r.RoleName).Count(&count)
	if count > 0 {
		return errno.ErrRoleExisting
	}
	if r.RoleID == "" {
		r.RoleID = tools.UUID()
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

func FindRole(roleID string) (*Role, error) {
	var role Role
	if err := db.Where(map[string]interface{}{"role_id": roleID}).
		Take(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ErrRoleNotFound
		}
		return nil, err
	}
	return &role, nil
}
