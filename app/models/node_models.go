package models

import "time"

type Node struct {
	ID           int       `form:"id"`
	Pid          int       `form:"pid" gorm:"type:int comment '上级ID';DEFAULT:0;"`
	IsType       int8      `form:"istype" gorm:"type:tinyint(1) comment '是否菜单:0菜单,1按钮';DEFAULT:0;"`
	Name         string    `form:"name" gorm:"type:varchar(50) comment '节点名称';not null;"`
	Icon         string    `form:"icon" gorm:"type:varchar(50) comment '图标'"`
	Url          string    `form:"url" gorm:"type:varchar(255) comment '菜单url'"`
	RouteName    string    `form:"routename" gorm:"type:varchar(100) comment '权限标识'"`
	Reorder      int8      `form:"reorder" gorm:"type:tinyint(1) comment '排序';DEFAULT:0;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Child []*Node                `gorm:"-"` //包含子菜单
	Role_is_active bool          `gorm:"-"` //某个角色是否拥有这个节点权限

}
