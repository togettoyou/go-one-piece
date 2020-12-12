package model

type CasbinRoleApiInfo struct {
	Path   string `json:"path" binding:"required" example:"/*"`
	Method string `json:"method" binding:"required,oneof=POST GET PATCH PUT DELETE *" example:"*"`
	ApiDes string `json:"api_des" binding:"omitempty"`
}

// CasbinUserRole 记录用户角色
type CasbinUserRole struct {
	PType    string `json:"-" gorm:"column:p_type;default:g"`
	Username string `json:"username" gorm:"column:v0"`
	RoleID   string `json:"role_id" gorm:"column:v1"`
	RoleName string `json:"role_name" gorm:"column:v2"`
}

// CasbinRoleApi 记录角色权限
type CasbinRoleApi struct {
	PType  string `json:"-" gorm:"column:p_type;default:p"`
	RoleID string `json:"role_id" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
	ApiDes string `json:"api_des" gorm:"column:v3"`
}
