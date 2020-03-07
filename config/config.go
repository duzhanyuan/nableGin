package config

import (

	"errors"
	"flag"
	"fmt"

	"gopkg.in/yaml.v2"
	"io/ioutil"

)

var (
	Conf        *conf        // 静态配置
	DynamicConf *dynamicConf // 动态配置

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

// 动态配置，程序运行过程中，可以动态更改的参数配置
type dynamicConf struct {
	UserDefaultName string `yaml:"user_default_name"`
	UserDefaultAge  int    `yaml:"user_default_age"`
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

	fmt.Printf("static config => [%#v]\n", Conf)

	DynamicConf = &dynamicConf{}
	if err := yaml.Unmarshal(content, DynamicConf); err != nil {
		return err
	}

	fmt.Printf("dynamic config => [%#v]\n", DynamicConf)

	return nil
}

