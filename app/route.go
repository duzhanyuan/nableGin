package app

import (
	. "nable.gin/config"
	"github.com/casbin/casbin/v2"
	redisadapter "github.com/casbin/redis-adapter/v2"
	middleware "nable.gin/app/middleware"
	"github.com/gin-gonic/gin"
	"nable.gin/app/controllers"
	"net/http"
)

//路由设置
func registerRouter(router *gin.Engine) {

	//CASBIN权限配置
	//配置casbin权限数据同步到REDIS
	adp := redisadapter.NewAdapter("tcp", Conf.RedisAddr)
	var Enforcer, _ = casbin.NewEnforcer("casbin.conf", adp)
	//设置初始权限 默认ADMIN用户拥有所有权限
	Enforcer.LoadPolicy()
	Enforcer.AddPolicy("anonymous", "/admin", "GET")
	Enforcer.AddPolicy("anonymous", "/admin/login/", "*")
	Enforcer.AddPolicy("admin", "/*", "*")
	if err := Enforcer.SavePolicy(); err != nil {
		panic(err)
	}

	//自定义404
	router.NoRoute(err404)

	//当前状态
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong",})
	})

	//欢迎页
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "welcome.html", gin.H{})
	})

	//获取验证码
	router.GET("/captcha", controllers.Captcha)

	//后台入口
	router.GET("/admin", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
	})

	//路由组设置
	admin := router.Group("/admin")
	admin.Use(middleware.NewAuthorizer(Enforcer))//使用权限验证中间件
	{
		//dash面板首页
		admin.GET("/dash/index", controllers.Index)

		//退出登陆 /admin/logout
		admin.GET("/logout", controllers.Logout)

		// 嵌套路由组 登陆控制层 /admin/login
		login := admin.Group("login")
		login.GET("/", controllers.GetLogin)
		login.POST("/", controllers.PostLogin)



	}



}


func err404(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
}

