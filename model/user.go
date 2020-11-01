package model

import (
	"github.com/satori/go.uuid"
	"go-one-server/util/errno"
	"gorm.io/gorm"
)

// 用户表
type User struct {
	Model
	UUID      uuid.UUID `json:"uuid" gorm:"size:36;not null;unique;comment:用户UUID"`
	Username  string    `json:"username" gorm:"size:16;not null;unique;comment:用户登录名"`
	Password  string    `json:"-"  gorm:"type:varchar(32);not null;comment:用户登录密码"`
	Salt      string    `json:"-" gorm:"type:varchar(12);not null;comment:混淆盐"`
	NickName  string    `json:"nick_name" gorm:"size:16;default:匿名;comment:用户昵称" `
	HeaderImg string    `json:"header_img" gorm:"default:https://avatars0.githubusercontent.com/u/55381228;comment:用户头像"`
}

// 创建钩子 https://gorm.io/zh_CN/docs/hooks.html
// 创建记录时将调用这些钩子方法
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.Must(uuid.NewV4(), nil)
	return nil
}

// 创建单条记录
func (u *User) Create() error {
	return db.Create(u).Error
}

// 软删除
func DeleteUser(uuid string) error {
	return db.Where("uuid = ?", uuid).Delete(&User{}).Error
}

// 根据uuid更新资料
func (u *User) UpdateUserInfo() error {
	// NickName和HeaderImg中只会更新非零值的字段
	return db.Model(&User{}).Where("uuid = ?", u.UUID).
		Updates(User{NickName: u.NickName, HeaderImg: u.HeaderImg}).
		Error
}

// 根据uuid更改密码
func (u *User) UpdateUserPw() error {
	// Select更新选定字段
	return db.Model(&User{}).Where("uuid = ?", u.UUID).
		Select("Password", "Salt").
		Updates(u).
		Error
}

// 根据uuid查询记录
func FindUser(uuid string) (*User, error) {
	var user User
	if err := db.Where(map[string]interface{}{"uuid": uuid}).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
