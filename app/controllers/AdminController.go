package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"nable.gin/libraries/viewdata"
	"net/http"
)

type AdminController struct {
}

func (admin *AdminController) Router(engine *gin.Engine) {
	engine.GET("/admin/dash/index", admin.Index)
	engine.GET("/admin/logout", admin.Logout)
}


//登陆页 admin.Login
func (admin *AdminController) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	v := viewdata.Default(ctx)

	////从闪存中获取消息
	flashes := session.Flashes()
	v.Set("flashes", flashes)
	session.Save()//移除闪存数据


	v.HTML(http.StatusOK, "admin_index.html")

}

/**
 * 后台首页面板
 * 请求类型：Get
 * 请求url：admin/logout
 */
func (admin *AdminController) Logout(ctx *gin.Context)  {

	session := sessions.Default(ctx)
	session.Delete("username")
	session.Delete("userid")
	session.AddFlash( "您已经成功退出系统")
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
	return


}

