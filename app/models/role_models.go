package models

import "time"

type Role struct {
	ID           int
	Name         string    `gorm:"type:varchar(20) comment '角色名称';not null;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time

}
