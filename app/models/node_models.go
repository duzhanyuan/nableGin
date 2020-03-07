package models

import "time"

type Node struct {
	ID           int
	Pid          int       `gorm:"type:int comment '上级ID';DEFAULT:0;"`
	IsType       int8      `gorm:"type:tinyint(1) comment '是否菜单:0菜单,1按钮';DEFAULT:0;"`
	Name         string    `gorm:"type:varchar(50) comment '节点名称';not null;"`
	Icon         string    `gorm:"type:varchar(50) comment '图标'"`
	Url          string    `gorm:"type:varchar(255) comment '菜单url'"`
	RouteName    string    `gorm:"type:varchar(100) comment '权限标识'"`
	Reorder      int8      `gorm:"type:tinyint(1) comment '排序';DEFAULT:0;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Child []*Node          `gorm:"-"`
}
