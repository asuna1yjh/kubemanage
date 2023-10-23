package models

// 角色菜单表
type SysRoleMenu struct {
	ID     int `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	RoleID int `json:"role_id" gorm:"column:role_id;type:int(11);not null;comment:角色ID"`
	MenuID int `json:"menu_id" gorm:"column:menu_id;type:int(11);not null;comment:菜单ID"`
}

func (s *SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (s *SysRoleMenu) TableComment() string {
	return "角色菜单表"
}
