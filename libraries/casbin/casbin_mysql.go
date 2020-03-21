package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
	. "nable.gin/config"
)
//
////
//////新增权限
////func Add(name string, path string, method string) bool {
////	e := initCasbin()
////	return e.AddPolicy(name, path, method)
////}
////
//////新增权限
////func AddGroup(name string, path string, method string) bool {
////	e := initCasbin()
////	return e.AddGroupingPolicy(name, path, method)
////}
//
////持久化到数据库
func InitCasbinMysql() *casbin.Enforcer {

	a, err := gormadapter.NewAdapter("mysql", Conf.MysqlDsn, true) //你的驱动和数据源
	if err != nil {
		log.Fatalf("gormadapter.NewAdapter error:%v", err)
	}
	e, _ := casbin.NewEnforcer("casbin.conf", a)


	e.LoadPolicy()
	return e
}
//
//
//
//// Check the permission.
////e.Enforce("alice", "data1", "read")
//
//// Modify the policy.
//// e.AddPolicy(...)
//// e.RemovePolicy(...)
//
//// Save the policy back to DB.
////e.SavePolicy()
