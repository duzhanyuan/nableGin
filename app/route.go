package app

import (
	"github.com/gin-gonic/gin"
	"nable.gin/app/controllers"
	"net/http"
	middleware "nable.gin/app/middleware"
	"nable.gin/libraries/casbin"
)

//路由设置
func registerRouter(router *gin.Engine) {

	//CASBIN权限配置
	//配置casbin权限数据同步到MYSQL || REDIS
	e := casbin.InitCasbinMysql()
	//设置初始权限 默认ADMIN用户拥有所有权限
	//e.LoadPolicy()
	//e.AddPolicy("anonymous", "/admin", "GET")
	//e.AddPolicy("anonymous", "/admin/login/", "GET")
	//e.AddPolicy("anonymous", "/admin/login/", "POST")
	//e.AddPolicy("admin", "/*", "*")
	//if err := e.SavePolicy(); err != nil {
	//	panic(err)
	//}

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
		ctx.Redirect(http.StatusSeeOther, "/admin/login/")//跳转
	})

	//路由组设置
	admin := router.Group("/admin")
	admin.Use(middleware.NewAuthorizer(e))//使用权限验证中间件
	{
		//dash面板首页
		admin.GET("/dash/index", controllers.Index)

		//退出登陆 /admin/logout
		admin.GET("/dash/logout", controllers.Logout)

		// 嵌套路由组
		// 登陆控制层 /admin/login
		login := admin.Group("login")
		login.GET("/", controllers.GetLogin)
		login.POST("/", controllers.PostLogin)

		// 用户管理控制层 /admin/user
		user := admin.Group("user")
		user.GET("/user", controllers.User)//列表
		user.GET("/add", controllers.Add)//新增用户
		user.POST("/store", controllers.Store)//提交新增用户
		user.POST("/del", controllers.Del)//软删除用户
		user.POST("/status", controllers.Status)//禁用/启用 用户状态切换
		user.GET("/edit/:id", controllers.EditBy)//用户编辑界面
		user.POST("/put", controllers.EditPut)//编辑提交

		//权限管理 /admin/node
		node := admin.Group("node")
		node.GET("/role", controllers.Role)//角色列表
		node.POST("/roleadd", controllers.Roleadd)//提交新增角色
		node.POST("/roleupdate", controllers.Roleupdate)//提交修改角色
		node.GET("/roledel", controllers.Roledel)//删除角色

		node.GET("/node", controllers.Node)//节点列表
		node.POST("/nodeadd", controllers.Nodeadd)//提交新增节点
		node.POST("/nodeupdate", controllers.Nodeupdate)//提交修改节点
		node.GET("/nodedel", controllers.Nodedel)//删除节点

		node.GET("/roleset/:id", controllers.Roleset)//角色设置权限列表页
		node.POST("/rolechecked", controllers.Rolechecked)//角色设置权限
		node.POST("/roleunchecked", controllers.Roleunchecked)//角色删除权限



	}



}


func err404(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
}

