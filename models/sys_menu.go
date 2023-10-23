package models

// 菜单表
type SysMenu struct {
	ID       int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name     string `json:"name" gorm:"column:name;type:varchar(100) COLLATE utf8mb4_unicode_ci;not null;default:'';comment:菜单名称;"`
	MenuCode string `json:"menu_code" gorm:"column:menu_code;type:varchar(100);not null;default:'';comment:菜单编码;"`
	ParentID int    `json:"parent_id" gorm:"column:parent_id;type:int(11);default:null;comment:父级菜单ID;"`
	NodeType int    `json:"node_type" gorm:"column:node_type;type:tinyint(4);default:1;comment:节点类型，1文件夹，2页面，3按钮；"`
	IconUrl  string `json:"icon_url" gorm:"column:icon_url;type:varchar(255) COLLATE utf8mb4_unicode_ci;not null;default:'';comment:菜单图标;"`
	Sort     int    `json:"sort" gorm:"column:sort;type:int(11);comment:排序;not null;default:1;"`
	LinkUrl  string `json:"link_url" gorm:"column:link_url;type:varchar(255);default:'';comment:菜单链接;"`
	Level    int    `json:"level" gorm:"column:level;type:int(11);not null;default:0;comment:菜单层级;"`
	Path     string `json:"path" gorm:"column:path;type:varchar(255);default:'';comment:树id的路径 整个层次上的路径id，逗号分隔，想要找父节点特别快;"`
	IsDelete int    `json:"is_delete" gorm:"column:is_delete;type:tinyint(1);not null;default:0;comment:'是否删除 1：已删除；0：未删除'"`
}
