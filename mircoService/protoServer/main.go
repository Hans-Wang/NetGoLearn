package main

import (
	"context"
	"fmt"
	proto "netgo/mircoService/protoServer/proto"
	micro "github.com/micro/go-micro/v2"

)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context,req *proto.HelloRequest, rsp *proto.HelloResponse) error{
	rsp.Greeting="你好，"+req.Name
	return nil
}

func main(){
	//创建服务，除了创建名称，还可以加其他类型的meta信息
	service:=micro.NewService(
		micro.Name("greeter.service"),
		micro.Version("latest"),
	)
	service.Init()
	//使用proto的注册Hanlder接口替换原来的默认Handler

	proto.RegisterGreeterHandler(service.Server(),new(Greeter))
	if err:=service.Run();err!=nil{
		fmt.Println(err)
	}
}