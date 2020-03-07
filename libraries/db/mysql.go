package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var mysqlDB *gorm.DB

// 初始化mysql
func StartMysql(dsn string, maxIdle, maxOpen int) (err error){
	mysqlDB, err = gorm.Open("mysql", dsn)

	//设置表前缀
	gorm.DefaultTableNameHandler = func(mysqlDB *gorm.DB, defaultTableName string) string {
		return "na_" + defaultTableName
	}

	if err == nil {
		mysqlDB.DB().SetMaxIdleConns(maxIdle)
		mysqlDB.DB().SetMaxOpenConns(maxOpen)
		mysqlDB.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
		//mysqlDB.SingularTable(true) 生成表名尾部不加S
		//mysqlDB.LogMode(true)// 启用Logger，显示SQL详细日志
	}

	return
}

// 获取mysql连接
func GetMysql() *gorm.DB {
	return mysqlDB
}

// 关闭mysql
func CloseMysql() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
}
