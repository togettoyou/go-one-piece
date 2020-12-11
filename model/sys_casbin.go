package model

type Casbin struct {
	PType  string `json:"p_type" gorm:"column:p_type"`
	RoleID string `json:"role_id" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}
