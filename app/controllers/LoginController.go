package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/rand"
	"nable.gin/app/models"
	"nable.gin/app/models/dto"
	"nable.gin/libraries/viewdata"
	"net/http"
	"time"
	"github.com/dchest/captcha"
	Dao "nable.gin/app/models/dao"
	Hash "nable.gin/libraries/helper"
)

type LoginController struct {
}

func (admin *LoginController) Router(engine *gin.Engine) {
	engine.GET("/admin/login", admin.GetLogin)
	engine.POST("/admin/login", admin.PostLogin)
	engine.GET("/admin/login/captcha", admin.Captcha)

}


//登陆页 admin.Login
func (admin *LoginController) GetLogin(ctx *gin.Context) {

	session := sessions.Default(ctx)
	v := viewdata.Default(ctx)

	////如果已经登陆
	userid := session.Get("userid")
	if userid != nil {
		session.AddFlash( "您已经登陆系统")
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/admin/dash/index")//跳转
		return
	}

	////从闪存中获取消息
	flashes := session.Flashes()
    v.Set("flashes", flashes)
	session.Save()//移除闪存数据

	//0-15随机数 随机背景图
	rand.Seed(time.Now().UnixNano())
	bg_num := rand.Intn(15)
	v.Set("bgnum", bg_num)

	v.HTML(http.StatusOK, "admin_login.html")


}



//func (admin *LoginController) GetLogin2(ctx *gin.Context) {
//
//	ctx.JSON(200, map[string]interface{}{"meeage": flashes,})
//




//登陆 admin.Login
func (admin *LoginController) PostLogin(ctx *gin.Context) {

	session := sessions.Default(ctx)

	//表单验证
	var params dto.Login_validate //登陆表单验证结构体 并返回表单值
	// 绑定DTO 进行表单验证
	if err := params.Bind(ctx); err != "" {
		session.AddFlash(err)
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
		return
	}

	//取出SESSION存储的验证码
	sessID := session.Get("CAPTCHA_Id")
	if sessID == nil {
		session.AddFlash( "验证码过期失效")
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
		return
	}

	if !captcha.VerifyString(sessID.(string), params.Capt_request) {
		session.AddFlash( "验证码错误")
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
		return
	}

	// 根据FORM params.Username获取用户
	var err_db error
	var userInfo models.User
	if userInfo, err_db = Dao.GetUserByName(params.Username); err_db != nil {
		session.AddFlash( "无效的用户名")
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
		return
	}

	//验证取出用户的密码是否匹配
	if ! Hash.CompareHashAndPassword(userInfo.Password,params.Password) {
		session.AddFlash( "无效的密码")
		session.Save()//消息放入闪存
		ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
		return
	}else{
		//密码匹配
		userInfo.Password="123456" //session中显示假密码
		session.Set("username", userInfo.Username)
		session.Set("userid", userInfo.ID)
		//session.Set("userInfo", map[string]interface{}{"userInfo": userInfo,})
		session.Save()
		ctx.Redirect(http.StatusSeeOther, "/admin/dash/index")//跳转
		return
	}




}




//验证码 admin.Captcha
func (admin *LoginController) Captcha(ctx *gin.Context) {

	//验证码ID
	captchaId := captcha.NewLen(4)
	//验证码ID存储到SESSION
	session := sessions.Default(ctx)
	session.Set("CAPTCHA_Id", captchaId)
	session.Save()

	ctx.Header("Content-Type", "image/png")
	if err := captcha.WriteImage(ctx.Writer, captchaId, 240, 80); err != nil {
		println(err)
	}


}

