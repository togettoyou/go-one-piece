package model

import (
	"go-one-server/util/errno"
	"gorm.io/gorm"
)

// Action 行为权限表
type Api struct {
	ID uint `json:"id" gorm:"primarykey"`
	ApiInfo
}

type ApiInfo struct {
	Path        string `json:"path" gorm:"comment:api路径" binding:"required" example:"/api/v1/api"`
	Method      string `json:"method" gorm:"type:varchar(6);comment:请求方法" binding:"required,oneof=POST GET PATCH PUT DELETE" example:"POST"`
	ApiGroup    string `json:"api_group" gorm:"default:base;comment:api组" binding:"required" example:"base"`
	Description string `json:"description" gorm:"comment:api中文描述" binding:"required" example:"api中文描述"`
}

func (a *Api) Create() error {
	return db.Create(a).Error
}

func GetApiList(page, pageSize int) (data *PaginationQ, err error) {
	var apis []*Api
	var total int64
	err = db.Model(&Api{}).Scopes(Count(&total)).Scopes(Paginate(&page, &pageSize)).Find(&apis).Error
	if err != nil {
		return nil, err
	}
	return &PaginationQ{
		PageSize: pageSize,
		Page:     page,
		Data:     apis,
		Total:    total,
	}, nil
}

func FindApiByID(id uint) (*Api, error) {
	var api Api
	if err := db.Where("id = ?", id).
		Take(&api).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ErrApiNotFound
		}
		return nil, err
	}
	return &api, nil
}

func FindApiInID(ids []uint) ([]Api, error) {
	apis := make([]Api, 0)
	if err := db.Where("id in (?)", ids).
		Find(&apis).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return apis, nil
}

func DelApi(id uint) error {
	return db.Where("id = ?", id).
		Delete(&Api{}).Error
}
