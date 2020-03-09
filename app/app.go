package app

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	. "nable.gin/config"
	"nable.gin/libraries/csrf"
	"nable.gin/libraries/helper"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)


func RunGin(port int) {

	//启动日志
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	//gin.DisableConsoleColor()
	//f, _ := os.Create(Conf.LogPath)// 记录到文件。
	//gin.DefaultWriter = io.MultiWriter(f) //记录到文件
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	app := gin.Default()
	//注册session，redis存储
	store, _ := redis.NewStore(10, "tcp", Conf.RedisAddr, "", []byte("secret"))
	store.Options(sessions.Options{MaxAge: 3600})
	app.Use(sessions.Sessions("GinSession", store))


	//全局使用CSRF中间件,权限中间件
	app.Use(csrf.Middleware())

	//视图模板自定义处理函数
	app.SetFuncMap(template.FuncMap{
		"strstr": helper.Strstr,//字符串比较
		"add": helper.Add,//数字相加
		"substr": helper.Substr,//字符串截取
		"date": helper.Date,//日期格式化
	})

	//注册静态资源
	app.Static("/static", "public/static/admin")
	app.LoadHTMLGlob("app/views/**/*")

	//注册路由
	registerRouter(app)

	srv := &http.Server{
		Addr:    ":"+strconv.Itoa(port),
		Handler: app,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")



}

