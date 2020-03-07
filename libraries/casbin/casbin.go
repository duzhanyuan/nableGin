package casbin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/casbin/casbin/v2"

)


func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}


func (c *Casbin) ServeHTTP() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		username := session.Get("username")
		if username == nil {
			username = "anonymous"
		}

		if !c.Check(ctx.Request, username.(string)) {
			session.AddFlash( "请先登陆系统")
			session.Save()//消息放入闪存
			ctx.Redirect(http.StatusSeeOther, "/admin/login")//跳转
			return
		}

		ctx.Next()
	}
}


type Casbin struct {
	enforcer *casbin.Enforcer
}


func (c *Casbin) Check(r *http.Request, username string) bool {
	method := r.Method
	path := r.URL.Path
	ok, _ := c.enforcer.Enforce(username, path, method)
	return ok
}

