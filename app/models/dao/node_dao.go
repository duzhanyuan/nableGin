package dao

import (
	"nable.gin/app/models"
	"nable.gin/libraries/db"
)

// 查询所有角色
func QueryNode() (Nodes []models.Node, err error) {
	res := db.GetMysql().Where("id > ?", 0).Order("reorder desc,id asc").Find(&Nodes)
	err = res.Error
	return
}

// 查询所有菜单
func QueryMenus() (Menus []models.NodeMenus, err error) {
	var nodes []models.Node
	res := db.GetMysql().
		Select("id,name").
		Where("is_type = ?", 0).
		Find(&nodes).Scan(&Menus)
	err = res.Error
	return
}


//根据模型节点权限新增入库
func NodeInsert(node models.Node) (err error) {
	res := db.GetMysql().Create(&node)
	err = res.Error
	return
}


//根据模型 编辑入库
func NodeUpdate(node models.Node) (err error) {

	//以下为不修改字段
	if node.ID != 0 {
		var oldData models.Node
		db.GetMysql().First(&oldData, node.ID)

		node.ID = oldData.ID
		node.CreatedAt = oldData.CreatedAt
	}


	res := db.GetMysql().Save(&node)
	err = res.Error
	return
}

//根据ID删除节点
func NodeDel(NodeId int) (err error) {
	var node models.Node
	res := db.GetMysql().Where("id = ?", NodeId).Delete(&node)
	err = res.Error
	return
}









