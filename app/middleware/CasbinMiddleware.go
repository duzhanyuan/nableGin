package casbin

import (
	"github.com/gin-contrib/sessions"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {

		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			username = "anonymous"
		}

		if !a.CheckPermission(c.Request,username.(string)) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request,username string) bool {
	user := username
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		panic(err)
	}

	return allowed
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(http.StatusNotFound, "403.html", gin.H{})
	c.AbortWithStatus(403)
}
