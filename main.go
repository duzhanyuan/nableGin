package main

import (
	"flag"
	"fmt"
	"os"

	. "nable.gin/config"
	"nable.gin/libraries/db"
	"nable.gin/app"
)

func main() {
	// 初始化配置文件
	flag.Parse()
	fmt.Print("InitConfig...\r")
	checkErr("InitConfig", InitConfig())
	fmt.Print("InitConfig Success!!!\n")

	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("StartMysql...\r")
	checkErr("StartMysql", db.StartMysql(Conf.MysqlDsn, Conf.MysqlMaxIdle, Conf.MysqlMaxOpen))
	fmt.Print("StartMysql Success!!!\n")

	//// 启动redis
	defer db.CloseRedis()
	fmt.Print("StartRedis...\r")
	checkErr("StartRedis", db.StartRedis(Conf.RedisAddr, Conf.RedisDB, Conf.RedisMaxIdle, Conf.RedisMaxOpen))
	fmt.Print("StartRedis Success!!!\n")

	//初始化数据库
	fmt.Print("DB Init...\r")
	app.Migrations()
	fmt.Print("DB Init Success!!!\n")

	// 开始运行iris框架
	fmt.Print("RunIris...\r")
	app.RunGin(Conf.ServerPort)

}

// 检查错误
func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf("%s Error: %v\n", errMsg, err)
		os.Exit(1)
	}
}
