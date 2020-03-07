package dao

import (
	"nable.gin/app/models"
	"nable.gin/libraries/db"
)

// 根据用户ID查询
func QueryUsersById(userId int) (userinfo models.User, err error) {
	res := db.GetMysql().Where("id = ?", userId).Find(&userinfo)
	err = res.Error
	return
}

// 根据用户名查询
func GetUserByName(userName string) (users models.User, err error) {
	res := db.GetMysql().Where("username = ?", userName).Find(&users)
	err = res.Error
	return
}

// 根据EMAIL查询
func GetUserByEmail(userEmail string) (users models.User, err error) {
	res := db.GetMysql().Where("email = ?", userEmail).Find(&users)
	err = res.Error
	return
}


// 根据用户名模糊查询
func QueryUsersByName(userName string) (users []*models.User, err error) {
	res := db.GetMysql().Where("username like ?", "%" + userName + "%").Find(&users)
	err = res.Error

	return
}

//根据模型入库
func UserInsert(u models.User) (err error) {
	res := db.GetMysql().Create(&u)
	err = res.Error
	return
}

//根据ID删除用户
func UserDel(userId int) (err error) {
	var u models.User
	res := db.GetMysql().Where("id = ?", userId).Delete(&u)
	err = res.Error
	return
}

//根据ID切换用户
func UserStatus(userId int,active int) (err error) {

	var u models.User
	res := db.GetMysql().Model(&u).Where("id = ?", userId).Update("is_active", active)
	err = res.Error
	return
}

//根据模型 编辑入库
func UserUpdate(u models.User) (err error) {

	//以下为不修改字段
	if u.ID != 0 {
		var oldUser models.User
		db.GetMysql().First(&oldUser, u.ID)

		u.RoleId = oldUser.RoleId
		u.Username = oldUser.Username
		u.Email = oldUser.Email
		u.Email = oldUser.Email
		u.IsActive = oldUser.IsActive
		u.Avatar = oldUser.Avatar
		u.LastIp = oldUser.LastIp
		u.CreatedAt = oldUser.CreatedAt
	}

	//u.DeletedAt = nil

	res := db.GetMysql().Save(&u)
	err = res.Error
	return
}


