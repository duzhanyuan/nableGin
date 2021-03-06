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

	v.Set("Nodelist", tmpNodelist)//列表
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
	ctx.Redirect(http.StatusSeeOther, "/admin/node/node")//跳转
	return


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




/**
 * 节点删除
 * 请求类型：GET
 * 请求url：node/roleset
 */
func Roleset(ctx *gin.Context)  {

	v := viewdata.Default(ctx)

	//读取出所有的权限节点
	nodeList,_ := dao.QueryNode()
	//将节点按树形分类，递归迭代数组排列
	tmpNodelist := services.Iteration(nodeList, 0)

	//读取当前角色信息
	id, _ := strconv.Atoi(ctx.Param("id"))
	roleInfo, _ := dao.GetRoleById(id)

	//根据角色所拥有的菜单 重置所有菜单里的 Role_is_active 选中值
	for idx, v := range tmpNodelist {
		for _, k := range roleInfo.Node {
			if (v.ID == k.ID){
				tmpNodelist[idx].Role_is_active=true
			}
		}
	}


	v.Set("Nodelist", tmpNodelist)//列表
	v.Set("role", roleInfo)//角色信息

	v.HTML(http.StatusOK, "admin_nodeset.html")
}


/**
 * 角色设置权限
 * 请求类型：POST
 * 请求url：node/rolechecked
 */
func Rolechecked(ctx *gin.Context)  {

	role_id, _ := strconv.Atoi(ctx.PostForm("role_id"))
	node_id, _ := strconv.Atoi(ctx.PostForm("node_id"))


	var status int
	var result string
	if err_db := dao.RoleInsertNodeId(role_id,node_id); err_db != nil {
		status = 1
		result = "权限设置失败"
	}else{
		status = 0
		result = "权限设置成功"
	}

	data := map[string]interface{}{
		"status": status,
		"result": result,
	}


	ctx.JSONP(http.StatusOK, data)


}




/**
 * 角色设置权限
 * 请求类型：POST
 * 请求url：node/rolechecked
 */
func Roleunchecked(ctx *gin.Context)  {

	role_id, _ := strconv.Atoi(ctx.PostForm("role_id"))
	node_id, _ := strconv.Atoi(ctx.PostForm("node_id"))

	dao.RoleDelNodeId(role_id,node_id)

	data := map[string]interface{}{
		"status": 0,
		"result": "权限删除成功",
	}
	ctx.JSONP(http.StatusOK, data)


}


/**
 * 用户角色权限设置
 * 请求类型：GET
 * 请求url：node/user
 */
func UserNodeSet(ctx *gin.Context)  {


	v := viewdata.Default(ctx)
	page := ctx.DefaultQuery("page", "1")
	pageNum, _ := strconv.Atoi(page) //string2int

	//接收搜索key
	search := ctx.DefaultQuery("search", "")

	//获取db连接类
	db := db.GetMysql()
	//默认搜索所有id>0 where作用域问题 要在{}外定义才能被Paginator引用
	where := db.Where("id > ?", 0)
	//如果不为空
	if search != "" {
		where = db.Where("truename LIKE ?", "%"+search+"%")
	}

	//定义模型为切片数组
	var userList []models.User

	//分页类获取数据 &userList 返回切片数组列表  Paginator为分页相关结构体
	Paginator := pagination.Paging(&pagination.Param{
		DB:      where,
		Page:    pageNum,
		Limit:   10,
		OrderBy: []string{"id desc"},
	}, &userList)

	v.Set("Userlist", userList)
	v.Set("TotalPage", Paginator.TotalPage)
	v.Set("Page", Paginator.Page)
	v.HTML(http.StatusOK, "admin_usernodeset.html")


}

