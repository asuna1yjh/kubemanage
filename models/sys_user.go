package models

import "gorm.io/gorm"

// 用户表
type SysUser struct {
	ID       int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Mobile   string `json:"mobile" gorm:"column:mobile;type:varchar(20);not null;default:'';comment:手机号;index:idx_mobile"`
	Name     string `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:用户名;index:idx_name"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);default:'';not null;comment:密码"`
	IsDelete int    `json:"is_delete" gorm:"column:is_delete;type:tinyint(1);not null;default:0;comment:'是否删除 1：已删除；0：未删除'"`
	*gorm.Model
}

func (u *SysUser) TableName() string {
	return "sys_user"
}

func (u *SysUser) TableComment() string {
	return "用户表"
}
