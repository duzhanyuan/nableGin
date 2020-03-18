package dao

import (
	"nable.gin/app/models"
	"nable.gin/libraries/db"
)


//根据模型角色入库
func RoleInsert(role models.Role) (err error) {
	res := db.GetMysql().Create(&role)
	err = res.Error
	return
}

//根据模型 编辑入库
func RoleUpdate(role models.Role) (err error) {

	//以下为不修改字段
	if role.ID != 0 {
		var oldRole models.Role
		db.GetMysql().First(&oldRole, role.ID)

		role.ID = oldRole.ID
		role.CreatedAt = oldRole.CreatedAt
	}


	res := db.GetMysql().Save(&role)
	err = res.Error
	return
}


//根据ID删除角色
func RoleDel(RoleId int) (err error) {
	var role models.Role
	res := db.GetMysql().Where("id = ?", RoleId).Delete(&role)
	err = res.Error
	return
}


// 根据ID查询角色 并读取包含的菜单
func GetRoleById(Id int) (role models.Role, err error) {
	db.GetMysql().Where("id = ?", Id).Find(&role)
	res := db.GetMysql().Model(role).Related(&role.Node, "Node")
	err = res.Error
	return
}


// 根据ID查询角色 并插入中间表设置权限
func RoleInsertNodeId(roleId int,nodeId int) (err error) {

	var role models.Role
	db.GetMysql().Preload("Node").First(&role, "id = ?", roleId)
	res := db.GetMysql().Model(&role).Association("Node").Append(&models.Node{ID: nodeId})
	err = res.Error
	return
}


// 根据ID查询角色 删除中间表权限
func RoleDelNodeId(roleId int,nodeId int) (err error) {

	var role models.Role
	db.GetMysql().Preload("Node").First(&role, "id = ?", roleId)
	res := db.GetMysql().Model(&role).Association("Node").Delete(&models.Node{ID: nodeId})

	err = res.Error
	return
}



// 根据ID查询角色
//func GetNodeByRoleId(Id int) (rolenode models.RoleNode, err error) {
//	res := db.GetMysql().Where("id = ?", Id).Find(&rolenode)
//	err = res.Error
//	return
//}

//// 根据ID查询角色所有菜单
//func GetNodePersonMenus() (rolenode models.RoleNode, err error) {
//	res := db.GetMysql().Where("id = ?", Id).Find(&rolenode)
//	err = res.Error
//	return
//}











