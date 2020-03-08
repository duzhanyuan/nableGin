package viewdata

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"nable.gin/libraries/csrf"
)

type ViewData map[string]interface{}

const DefaultKey = "viewdata.data"

func Default(ctx *gin.Context) (v ViewData) {
	val, exists := ctx.Get(DefaultKey)
	if exists {
		v = val.(ViewData)
	} else {
		v = ViewData{"ctx": ctx}
		ctx.Set(DefaultKey, v)
	}
	return
}

func (v ViewData) Set(key string, data interface{}) {
	v[key] = data
}

func (v ViewData) Get(key string) (interface{}, bool) {
	data, ok := v[key]
	return data, ok
}

func (v ViewData) GetContext() (ctx *gin.Context) {
	c, _ := v.Get("ctx")
	ctx, _ = c.(*gin.Context)
	return
}

func (v ViewData) GetInt(key string) (i int, b bool) {
	data, b := v.Get(key)
	if b {
		i, b = data.(int)
	}
	return
}

func (v ViewData) GetIntDefault(key string, def int) int {
	if d, b := v.GetInt(key); b {
		return d
	}
	return def
}

func (v ViewData) GetFloat64(key string) (i float64, b bool) {
	data, b := v.Get(key)
	if b {
		i, b = data.(float64)
	}
	return
}

func (v ViewData) GetFloat64Default(key string, def float64) float64 {
	if d, b := v.GetFloat64(key); b {
		return d
	}
	return def
}

func (v ViewData) GetString(key string) (i string, b bool) {
	data, b := v.Get(key)
	if b {
		i, b = data.(string)
	}
	return
}

func (v ViewData) GetStringDefault(key string, def string) string {
	if d, b := v.GetString(key); b {
		return d
	}
	return def
}

func (v ViewData) HTML(code int, name string) {
	v.Set("username", sessions.Default(v.GetContext()).Get("username"))
	v.Set("userid", sessions.Default(v.GetContext()).Get("userid"))
	v.Set("Token", csrf.GetToken(v.GetContext()))
	v.Set("AdminTitle", "nAble Pro")
	v.Set("path", v.GetContext().FullPath())
	v.GetContext().HTML(code, name, v)
}

func (v ViewData) Clear() {
	ctx := v.GetContext()
	for k := range v {
		delete(v, k)
	}
	v["ctx"] = ctx
}
