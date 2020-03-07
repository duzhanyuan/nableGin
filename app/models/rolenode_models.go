package models

type RoleNode struct {
	ID              int
	RoleId          int       `gorm:"type:int comment '角色ID';DEFAULT:0;"`
	NodeId          int       `gorm:"type:int comment '节点ID';DEFAULT:0;"`

}
