package casbin

import (
	"github.com/casbin/casbin/v2"
	redisadapter "github.com/casbin/redis-adapter/v2"
	. "nable.gin/config"
)


//持久化到数据库
func InitCasbinRedis() *casbin.Enforcer {
	a := redisadapter.NewAdapter("tcp",  Conf.RedisAddr)
	e , _ := casbin.NewEnforcer("casbin.conf", a)
	e.LoadPolicy()
	return e
}

