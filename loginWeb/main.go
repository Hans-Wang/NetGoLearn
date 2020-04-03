package main

import (
	"fmt"
	"netgo/loginWeb/dao"
	"netgo/loginWeb/routers"
)



func main() {
	err := dao.InitMysql()
	if err != nil {
		fmt.Printf("init MySQL failed, err :%v\n", err)
		return
	}

	r := routers.SetupRouter() //初始化

	r.Run()
}
