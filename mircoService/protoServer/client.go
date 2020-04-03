package main

import (
	"context"
	"fmt"
	proto "netgo/mircoService/protoServer/proto"
	"github.com/micro/go-micro/v2"
)

//
//请问Proto的MicroService方式
//如何增加一个Handler Hello2，保证他们都可以被请求？
func main(){
	service:=micro.NewService()
	service.Init()
	//创建专属于Greeter服务的配套客户端
	greeter:=proto.NewGreeterService("greeter.service",service.Client())
	//直接调用Greeter服务的Hello函数
	rsp,err:=greeter.Hello(context.TODO(),&proto.HelloRequest{Name:"小明"})
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}