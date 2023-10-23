package models

// 角色表
type SysRole struct {
	ID       int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Code     string `json:"code" gorm:"column:code;type:varchar(100);not null;default:'';comment:角色编码;index:idx_code"`
	Name     string `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:角色名称;index:idx_name"`
	IsDelete int    `json:"is_delete" gorm:"column:is_delete;type:tinyint(1);not null;default:0;comment:'是否删除 1：已删除；0：未删除'"`
}

func (s *SysRole) TableName() string {
	return "sys_role"
}

func (s *SysRole) TableComment() string {
	return "角色表"
}
