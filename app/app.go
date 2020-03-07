package app

import (
	. "nable.gin/config"
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
	"nable.gin/libraries/csrf"
	"nable.gin/libraries/helper"
	"net/http"
	"os"
	"io"
	"os/signal"
	"strconv"
	"time"
	"html/template"
	"github.com/casbin/casbin/v2"
	redisadapter "github.com/casbin/redis-adapter/v2"
	cm "nable.gin/libraries/casbin"

)



func RunGin(port int) {

	//启动日志
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	//gin.DisableConsoleColor()
	f, _ := os.Create(Conf.LogPath)// 记录到文件。
	//gin.DefaultWriter = io.MultiWriter(f) //记录到文件
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)


	app := gin.Default()
	//注册session，redis存储
	store, _ := redis.NewStore(10, "tcp", Conf.RedisAddr, "", []byte("secret"))
	app.Use(sessions.Sessions("GinSession", store))

	//CASBIN权限配置
	//配置casbin权限数据同步到REDIS
	adp := redisadapter.NewAdapter("tcp", Conf.RedisAddr)
	var Enforcer, _ = casbin.NewEnforcer("casbin.conf", adp)
	//设置初始权限 默认ADMIN用户拥有所有权限
	Enforcer.LoadPolicy()
	Enforcer.AddPolicy("anonymous", "/", "GET")
	Enforcer.AddPolicy("anonymous", "/static/*", "GET")
	Enforcer.AddPolicy("anonymous", "/admin", "GET")
	Enforcer.AddPolicy("anonymous", "/admin/login", "GET")
	Enforcer.AddPolicy("anonymous", "/admin/login/captcha", "GET")
	Enforcer.AddPolicy("anonymous", "/admin/login", "POST")
	Enforcer.AddPolicy("admin", "/*", "*")
	if err := Enforcer.SavePolicy(); err != nil {
		panic(err)
	}
	//设置权限中间件
	casbinMiddleware := cm.New(Enforcer)

	//全局使用CSRF中间件,权限中间件
	app.Use(csrf.Middleware(),casbinMiddleware.ServeHTTP())

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

