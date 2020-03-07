package config

import (

	"errors"
	"flag"
	"fmt"
	"nable.gin/libraries/helper"

	"gopkg.in/yaml.v2"
	"io/ioutil"

)

var (
	Conf        *conf        // 静态配置
	_path       string

)

const SERVER_NAME = "config"

// 静态配置，程序启动后无法再做更改的参数配置
type conf struct {
	ServerPort   int    `yaml:"server_port"`
	LogPath      string `yaml:"log_path"`
	MysqlDsn     string `yaml:"mysql_dsn"`
	MysqlMaxIdle int    `yaml:"mysql_max_idle"`
	MysqlMaxOpen int    `yaml:"mysql_max_open"`
	RedisAddr    string `yaml:"redis_addr"`
	RedisDB      int    `yaml:"redis_db"`
	RedisMaxIdle int    `yaml:"redis_max_idle"`
	RedisMaxOpen int    `yaml:"redis_max_open"`
	AdminTitle string   `yaml:"admin_title"`
}


// 初始化解析参数
func init() {
	flag.StringVar(&_path, "c", SERVER_NAME + ".yml", "default config path")
}

// 从配置文件中加载配置
func InitConfig() error {
	var err error
	var content []byte

	content, err = ioutil.ReadFile(_path)

	if err != nil {
		return err
	}

	if len(content) == 0 {
		return errors.New("not found nothing config")
	}

	Conf = &conf{}
	if err := yaml.Unmarshal(content, Conf); err != nil {
		return err
	}


	timedate := helper.DateYmdFormat()
	Conf.LogPath = string(Conf.LogPath)+"gin_"+string(timedate)+".log"

	fmt.Printf("static config => [%#v]\n", Conf)


	return nil
}

