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
// 钩子中的db操作需使用事务tx
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var count int64
	tx.Model(&User{}).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		return errno.ErrUserExisting
	}
	u.UUID = uuid.Must(uuid.NewV4(), nil)
	return nil
}

// 创建单条记录
func (u *User) Create() error {
	return db.Create(u).Error
}

// 软删除
func DeleteUser(username string) error {
	return db.Where("username = ?", username).Delete(&User{}).Error
}

// 根据username更新资料
func (u *User) UpdateUserInfo() error {
	// NickName和HeaderImg中只会更新非零值的字段
	return db.Model(&User{}).Where("username = ?", u.Username).
		Updates(User{NickName: u.NickName, HeaderImg: u.HeaderImg}).
		Error
}

// 根据username更改密码
func (u *User) UpdateUserPw() error {
	// Select更新选定字段
	return db.Model(&User{}).Where("username = ?", u.Username).
		Select("Password", "Salt").
		Updates(u).
		Error
}

// 根据username查询记录
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

// 分页获取用户列表
func GetUserList(page, pageSize int) (data *PaginationQ, err error) {
	var users []*User
	var total int64
	err = db.Model(&User{}).Scopes(Count(&total)).Scopes(Paginate(page, pageSize)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &PaginationQ{
		PageSize: pageSize,
		Page:     page,
		Data:     users,
		Total:    total,
	}, nil
}
