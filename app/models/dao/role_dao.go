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







