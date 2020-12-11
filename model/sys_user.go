package model

import (
	"go-one-server/util/errno"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	Model
	Username string `json:"username" gorm:"size:20;not null;unique;comment:用户名"`
	Password string `json:"-"  gorm:"type:char(32);not null;comment:登录密码"`
	Salt     string `json:"-" gorm:"type:char(12);not null;comment:混淆盐"`
	Remark   string `json:"remark" gorm:"comment:备注"`
	RoleID   string `json:"role_id" gorm:"not null;type:varchar(32);comment:角色ID"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID;references:RoleID"`
}

func (u *User) Create() error {
	return db.Create(u).Error
}

func FindUser(username string) (*User, error) {
	var user User
	if err := db.Where(map[string]interface{}{"username": username}).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
