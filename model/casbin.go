package model

// CasbinUserRole 记录用户角色
type CasbinUserRole struct {
	PType    string `json:"-" gorm:"column:p_type;default:g"`
	Username string `json:"username" gorm:"column:v0"`
	RoleKey  string `json:"role_key" gorm:"column:v1"`
	RoleName string `json:"role_name" gorm:"column:v2"`
}

// CasbinRoleApi 记录角色权限
type CasbinRoleApi struct {
	PType   string `json:"-" gorm:"column:p_type;default:p"`
	RoleKey string `json:"-" gorm:"column:v0"`
	ApiID   string `json:"api_id" gorm:"column:v1"`
	Path    string `json:"path" gorm:"column:v2"`
	Method  string `json:"method" gorm:"column:v3"`
	ApiDes  string `json:"api_des" gorm:"column:v4"`
}
