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
}

// 创建钩子 https://gorm.io/zh_CN/docs/hooks.html
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var count int64
	tx.Model(&User{}).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		return errno.ErrUserExisting
	}
	return nil
}

func (u *User) Create() error {
	return db.Create(u).Error
}

// 根据username查询记录
func FindUser(username string) (*User, error) {
	var user User
	if err := db.Where(map[string]interface{}{"username": username}).
		Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
