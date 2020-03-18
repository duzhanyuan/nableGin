package models

import "time"

type Role struct {
	ID           int       `form:"id"`
	Name         string    `form:"name" gorm:"type:varchar(20) comment '角色名称';not null;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time

	Node []Node `json:"Node" gorm:"many2many:role_nodes;"`
}
