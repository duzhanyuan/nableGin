package app

import (
	"github.com/gin-gonic/gin"
	"nable.gin/app/controllers"
	"net/http"
)

//路由设置
func registerRouter(router *gin.Engine) {

	//当前状态
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong",})
	})

	//欢迎页
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{})
	})

	//登陆页控制层
	new(controllers.LoginController).Router(router)
	new(controllers.AdminController).Router(router)





}
