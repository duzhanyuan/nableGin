package models

import "time"

type User struct {
	ID           int       `form:"id"`
	RoleId       int       `gorm:"type:int comment '角色ID，默认3为普通用户';DEFAULT:3"`
	Truename     string    `form:"truename" gorm:"type:varchar(50) comment '姓名' "`
	Username     string    `form:"username" gorm:"type:varchar(12) comment '用户名';unique;not null;"`
	Password     string    `form:"password" gorm:"type:varchar(256) comment '密码'"`
	Sex          string    `form:"sex" gorm:"type:varchar(10) comment '性别';DEFAULT:'先生';"`
	Email        string    `form:"email" gorm:"type:varchar(100) comment '邮箱';unique;"`
	Mobile       string    `form:"mobile" gorm:"type:varchar(15) comment '手机号'"`
	IsActive     int8      `gorm:"type:tinyint(1) comment '是否激活';DEFAULT:1"`
	Avatar       string    `gorm:"type:varchar(100) comment '头像';DEFAULT:'/static/img/default.jpg';"`
	LastIp       string    `gorm:"type:varchar(15) comment 'IP'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time


}

