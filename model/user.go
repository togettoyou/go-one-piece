package model

import (
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
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

func initUser() {
	// 如果数据库中不存在该表，则创建表
	if !db.Migrator().HasTable(&User{}) {
		// 创建表时添加后缀
		if err := db.Set("gorm:table_options",
			"ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&User{}); err != nil {
			zap.L().Error(err.Error())
		}
	}
}

// 创建钩子 https://gorm.io/zh_CN/docs/hooks.html
// 创建记录时将调用这些钩子方法
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.Must(uuid.NewV4(), nil)
	return nil
}

// 创建单条记录
func (u *User) Create() error {
	return db.Create(&u).Error
}

type Users struct {
	Users []User `json:"users"`
}

// 批量插入记录
func (us *Users) Create() error {
	return db.Create(&us.Users).Error
}
