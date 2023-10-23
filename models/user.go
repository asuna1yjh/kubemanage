package models

import "gorm.io/gorm"

// Greeter is a Greeter model.
type User struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	UserID   int64  `gorm:"column:user_id;type:bigint(20);not null;default:0;comment:用户ID" json:"user_id"`
	UserName string `gorm:"column:username;type:varchar(255);not null;comment:用户名;unique" json:"username"`
	NickName string `gorm:"column:nickname;type:varchar(255);comment:昵称" json:"nickname"`
	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`
	Email    string `gorm:"column:email;type:varchar(255);comment:邮箱" json:"email"`
	Phone    string `gorm:"column:phone;type:varchar(255);comment:手机号" json:"phone"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
	RoleID   int64  `gorm:"column:role_id;type:bigint(20);comment:角色ID" json:"role_id"`
	gorm.Model
}

func (u *User) TableName() string {
	return "userinfo"
}
