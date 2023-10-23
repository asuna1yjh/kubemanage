package models

// 用户角色表
type SysUserRole struct {
	ID     int `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	UserID int `json:"user_id" gorm:"column:user_id;type:int(11);not null;comment:用户ID"`
	RoleID int `json:"role_id" gorm:"column:role_id;type:int(11);not null;comment:角色ID"`
}

func (s *SysUserRole) TableName() string {
	return "sys_user_role"
}

func (s *SysUserRole) TableComment() string {
	return "用户角色表"
}
