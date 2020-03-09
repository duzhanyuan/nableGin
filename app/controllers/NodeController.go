package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"nable.gin/app/models"
	"nable.gin/app/models/dao"
	"nable.gin/app/models/dto"
	"nable.gin/libraries/db"
	"nable.gin/libraries/pagination"
	"nable.gin/libraries/viewdata"
	"nable.gin/services"
	"net/http"
	"strconv"
)

/**
 * 用户列表页面板
 * 请求类型：Get
 * 请求url：admin/node/role
 */
func Role(ctx *gin.Context) {
	session := sessions.Default(ctx)
	v := viewdata.Default(ctx)

	////从闪存中获取消息
	flashes := session.Flashes()
	v.Set("flashes", flashes)
	session.Save() //移除闪存数据

	//接收页码 如果页码不存在 默认为1
	page := ctx.DefaultQuery("page", "1")
	pageNum, _ := strconv.Atoi(page) //string2int

	//获取db连接类
	db := db.GetMysql()
	//默认搜索所有id>0 where作用域问题 要在{}外定义才能被Paginator引用
	where := db.Where("id > ?", 0)

	//定义模型为切片数组
	var roleList []models.Role

	//分页类获取数据 &userList 返回切片数组列表  Paginator为分页相关结构体
	Paginator := pagination.Paging(&pagination.Param{
		DB:      where,
		Page:    pageNum,
		Limit:   10,
		OrderBy: []string{"id desc"},
	}, &roleList)


	v.Set("Rolelist", roleList)
	v.Set("TotalPage", Paginator.TotalPage)//分页总数
	v.Set("Page", Paginator.Page)//当前页码

	v.HTML(http.StatusOK, "admin_role.html")

}


/**
 * 角色添加
 * 请求类型：POST
 * 请求url：admin/node/roleadd
 */
func Roleadd(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	var params dto.Roleadd_validate //登陆表单验证结构体 并返回表单值

	// 绑定DTO 进行表单验证
	if err := params.Bind(ctx); err != "" {
		session.AddFlash(err)
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/node/role")//跳转
		return
	}

	//入库模型定义
	var roleForm models.Role
	if err := ctx.ShouldBind(&roleForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//入库
	if err_db := dao.RoleInsert(roleForm); err_db != nil {
		session.AddFlash("添加失败！")
	}else{
		session.AddFlash("添加成功！")
	}
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/admin/node/role") //跳转


}


/**
 * 角色修改
 * 请求类型：POST
 * 请求url：node/roleupdate
 */
func Roleupdate(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	var params dto.Roleadd_validate //登陆表单验证结构体 并返回表单值

	// 绑定DTO 进行表单验证
	if err := params.Bind(ctx); err != "" {
		session.AddFlash(err)
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/node/role")//跳转
		return
	}

	//入库模型定义
	var roleForm models.Role
	if err := ctx.ShouldBind(&roleForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//编辑入库
	err_db := dao.RoleUpdate(roleForm)
	if (err_db != nil) {
		session.AddFlash("修改失败！")
	}else{
		session.AddFlash("修改成功！")
	}
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/admin/node/role") //跳转

}



/**
 * 角色删除
 * 请求类型：Get
 * 请求url：node/roledel
 */
func Roledel(ctx *gin.Context)  {
	session := sessions.Default(ctx)

	id,_ := strconv.Atoi(ctx.Query("id"))
	if(id <= 3){
		session.AddFlash("系统保留角色！")
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/node/role") //跳转
		return
	}

	err_db := dao.RoleDel(id)
	result := "删除成功"
	if (err_db != nil) {
		result="删除失败"
	}

	session.AddFlash(result)
	session.Save()//消息放入闪存
	ctx.Redirect(http.StatusSeeOther, "/admin/node/role") //跳转



}



/**
 * 节点列表页面板
 * 请求类型：Get
 * 请求url：node/node
 */
func Node(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	v := viewdata.Default(ctx)

	////从闪存中获取消息
	flashes := session.Flashes()
	v.Set("flashes", flashes)
	session.Save() //移除闪存数据

	//所有角色
	nodeList,_ := dao.QueryNode()
	//将角色按树形分类，递归迭代数组排列
	tmpNodelist := services.Iteration(nodeList, 0)

	//取出所有菜单
	menuList,_ := dao.QueryMenus()

	v.Set("Nodelist", tmpNodelist)//用户列表
	v.Set("Menulist", menuList)//菜单列表
	v.HTML(http.StatusOK, "admin_node.html")

}


/**
 * 节点添加
 * 请求类型：POST
 * 请求url：node/nodeadd
 */
func Nodeadd(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	var params dto.Nodeadd_validate //登陆表单验证结构体 并返回表单值

	// 绑定DTO 进行表单验证
	if err := params.Bind(ctx); err != "" {
		session.AddFlash(err)
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/node/node")//跳转
		return
	}

	//入库模型定义
	var nodeForm models.Node
	if err := ctx.ShouldBind(&nodeForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//入库
	if err_db := dao.NodeInsert(nodeForm); err_db != nil {
		session.AddFlash("添加失败！")
	}else{
		session.AddFlash("添加成功！")
	}
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/admin/node/node") //跳转



}


/**
 * 节点编辑
 * 请求类型：POST
 * 请求url：node/nodeupdate
 */
func Nodeupdate(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	var params dto.Nodeadd_validate //登陆表单验证结构体 并返回表单值

	// 绑定DTO 进行表单验证
	if err := params.Bind(ctx); err != "" {
		session.AddFlash(err)
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/node/node")//跳转
		return
	}

	//入库模型定义
	var nodeForm models.Node
	if err := ctx.ShouldBind(&nodeForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//入库
	if err_db := dao.NodeUpdate(nodeForm); err_db != nil {
		session.AddFlash("编辑失败！")
	}else{
		session.AddFlash("编辑成功！")
	}
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/admin/node/node") //跳转


}



/**
 * 节点删除
 * 请求类型：GET
 * 请求url：node/nodedel
 */
func Nodedel(ctx *gin.Context)  {
	session := sessions.Default(ctx)

	id,_ := strconv.Atoi(ctx.Query("id"))
	if(id < 10){
		session.AddFlash("系统保留权限！")
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/node/node") //跳转
		return
	}

	err_db := dao.NodeDel(id)
	result := "删除成功"
	if (err_db != nil) {
		result="删除失败"
	}
	session.AddFlash(result)
	session.Save()//消息放入闪存
	ctx.Redirect(http.StatusSeeOther, "/admin/node/node") //跳转

}

