package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"nable.gin/app/models"
	"nable.gin/app/models/dto"
	"nable.gin/libraries/db"
	"nable.gin/libraries/pagination"
	"nable.gin/libraries/viewdata"
	"net/http"
	"strconv"
	Dao "nable.gin/app/models/dao"
	Hash "nable.gin/libraries/helper"
)



/**
 * 用户列表页面板
 * 请求类型：Get
 * 请求url：admin/user/user
 */
func User(ctx *gin.Context) {

	v := viewdata.Default(ctx)
	page := ctx.DefaultQuery("page", "1")
	pageNum, _ := strconv.Atoi(page)//string2int

	//接收搜索key
	search := ctx.DefaultPostForm("search", "")

	//获取db连接类
	db := db.GetMysql()
	//默认搜索所有id>0 where作用域问题 要在{}外定义才能被Paginator引用
	where := db.Where("id > ?", 0)
	//如果不为空
	if(search != ""){
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
	v.HTML(http.StatusOK, "admin_user.html")


}


/**
 * 用户增加页面板
 * 请求类型：Get
 * 请求url：admin/user/add
 */
func Add(ctx *gin.Context)  {
	session := sessions.Default(ctx)
	v := viewdata.Default(ctx)

	////从闪存中获取消息
	flashes := session.Flashes()
	v.Set("flashes", flashes)
	session.Save()//移除闪存数据

	v.HTML(http.StatusOK, "admin_useradd.html")

}



/**
 * 用户增加页面板
 * 请求类型：Post
 * 请求url：user/store
 */
func Store(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	var params dto.Useradd_validate //登陆表单验证结构体 并返回表单值

	// 绑定DTO 进行表单验证
	if err := params.Bind(ctx); err != "" {
		session.AddFlash(err)
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/admin/user/add")//跳转
		return
	}

	// 根据FORM params.Username获取用户 如果用户名已经注册 返回错误消息
	var err_db error
	var userInfo models.User
	userInfo, err_db = Dao.GetUserByName(params.Username)
	if (userInfo.Username == params.Username && err_db == nil) {
		session.AddFlash("用户名已经注册")
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/admin/user/add")//跳转
		return
	}

	// 根据FORM params.Email获取用户 如果邮箱已经注册 返回错误消息
	userInfo, err_db = Dao.GetUserByEmail(params.Email)
	if (userInfo.Email == params.Email && err_db == nil) {
		session.AddFlash("邮箱已经注册")
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/admin/user/add")//跳转
		return
	}

	//入库模型定义
	var userForm models.User
	if err := ctx.ShouldBind(&userForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userForm.Password,_ = Hash.PasswordEncrypt(userForm.Password)//加密

	//入库
	err_db = Dao.UserInsert(userForm)
	if (err_db != nil) {
		session.AddFlash("添加失败！")
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/admin/user/add")//跳转
	}else{
		session.AddFlash("添加成功！")
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/admin/user/add")//跳转
	}





}

/**
 * 用户删除
 * 请求类型：POST
 * 请求url：user/del
 */
func Del(ctx *gin.Context)  {

	id,_ := strconv.Atoi(ctx.PostForm("id"))
	err_db := Dao.UserDel(id)

	status := 0
	result := "删除成功"
	if (err_db != nil) {
		status=1
		result="删除失败"
	}

	data := map[string]interface{}{
		"status":  status,
		"result":  result,
	}
	ctx.JSONP(http.StatusOK, data)

}



/**
 * 用户状态切换
 * 请求类型：POST
 * 请求url：user/status
 */
func Status(ctx *gin.Context)  {

	id,_ := strconv.Atoi(ctx.PostForm("id"))
	active,_ := strconv.Atoi(ctx.PostForm("active"))
	status := 0
	result := "用户状态切换成功。"

	if(id <= 3){
		status=1
		result="系统保留用户，禁用失败。"
	}else{
		err_db := Dao.UserStatus(id,active)
		if (err_db != nil) {
			status=1
			result="用户状态切换失败。"
		}
	}

	data := map[string]interface{}{
		"status":  status,
		"result":  result,
	}
	ctx.JSONP(http.StatusOK, data)

}


/**
 * 用户编辑界面
 * 请求类型：GET
 * 请求url：user/edit/1
 */
func EditBy(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	v := viewdata.Default(ctx)
	id,_ := strconv.Atoi(ctx.Param("id"))

	////从闪存中获取消息
	flashes := session.Flashes()
	v.Set("flashes", flashes)
	session.Save()//移除闪存数据

	var err_db error
	var userInfo models.User
	userInfo, err_db = Dao.QueryUsersById(id)
	if (err_db != nil) {
		session.AddFlash( "编辑用户错误！")
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/user/user")//跳转
		return
	}

	v.Set("userinfo", userInfo)
	v.HTML(http.StatusOK, "admin_useredit.html")



}



/**
 * 用户编辑动作
 * 请求类型：POST
 * 请求url：user/put
 */
func EditPut(ctx *gin.Context)  {
	session := sessions.Default(ctx)

	var params dto.Useredit_validate //登陆表单验证结构体 并返回表单值

	userid := ctx.PostForm("id")
	// 绑定DTO 进行表单验证
	if err := params.Bind(ctx); err != "" {
		session.AddFlash( err)
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/user/edit/"+userid)//跳转
		return
	}

	//入库模型定义
	var userForm models.User
	if err := ctx.ShouldBind(&userForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userForm.Password,_ = Hash.PasswordEncrypt(userForm.Password)//加密


	//编辑入库
	err_db := Dao.UserUpdate(userForm)
	if (err_db != nil) {
		session.AddFlash("编辑失败！")
	}else{
		session.AddFlash("编辑成功！")
	}
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/admin/user/edit/"+userid)//跳转



}


