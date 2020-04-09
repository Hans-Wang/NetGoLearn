package main

import (
	"fmt"
	"go.uber.org/zap"
	"netgo/loginWeb/config"
	"netgo/loginWeb/dao"
	"netgo/loginWeb/logger"
	"netgo/loginWeb/routers"
	"os"
)

func main() {

	//用配置文件初始化
	if len(os.Args) < 2 {
		return
	}

	if err := config.Init(os.Args[1]); err != nil {
		fmt.Printf("config.Init failed, err:%v\n", err)
		return
	}

	//用json字符串初始化
	//	s:=`
	//{
	//  "server": {
	//    "port": 8080
	//  },
	//  "mysql": {
	//    "host": "47.104.241.166",
	//    "port": 3306,
	//    "db": "gin-blog",
	//    "username": "gopher",
	//    "password": "2020O229_"
	//  },
	//  "redis": {
	//    "host": "47.104.241.166",
	//    "port": 6379,
	//    "db": 0,
	//    "password": ""
	//  },
	//  "log": {
	//    "lever": "debug",
	//    "filename": "gin-blog.log",
	//    "maxsize": 500,
	//    "max_age": 7,
	//    "max_backups": 10
	//  }
	//}`
	//	if err := config.InitFromStr(s);err != nil {
	//		fmt.Printf("config.Init failed, err:%v\n", err)
	//		return
	//	}

	// 初始化logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:#{err}\n")
		return
	}
	// 初始化 mysql
	if err := dao.InitMysql(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init MySQL failed, err:#{err}\n")
		logger.Logger.Error("init MySQL failed", zap.Any("error", err))
		return
	}
	// 初始化 redis
	//if err := dao.InitRedis(config.Conf.RedisConfig);err != nil{
	//	fmt.Printf("init redis failed, err:#{err}\n")
	//	return
	//}

	logger.Logger.Debug("start project...")
	//err := dao.InitMysql()
	//if err != nil {
	//	fmt.Printf("init MySQL failed, err :%v\n", err)
	//	return
	//}

	r := routers.SetupRouter() //初始化

	r.Run()
}
