package config

import (
	"encoding/json"
	"io/ioutil"
)

// AppConfig 应用的配置结构体
type AppConfig struct {
	*ServerConfig `json:"server" ini:"server"`
	*MySQLConfig  `json:"mysql" ini:"mysql"`
	*RedisConfig  `json:"redis" ini:"redis"`
	*LogConfig    `json:"log" ini:"log"`
}

// ServerConfig 服务的配置
type ServerConfig struct {
	Port int `json:"port"`
}

// MySQLConfig mysql数据库配置
type MySQLConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DB       int    `json:"db"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

var Conf = new(AppConfig) //定义了全局的配置文件

func Init(file string) error {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, Conf); err != nil {
		return err
	}
	return nil
}

func InitFromStr(str string) error {
	if err := json.Unmarshal([]byte(str), Conf); err != nil {
		return err
	}
	return nil
}
