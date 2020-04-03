package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
)

type Greeter struct {

}

func (g *Greeter)Hello(ctx context.Context, name *string, msg *string) error {
	fmt.Println("A new Request Comes")
	*msg="你好，"+*name+"-----这是一条Response的回复"
	return nil
}

func (g *Greeter)Hello2(ctx context.Context, arg *string, msg *string) error {
	fmt.Println("hello22222的请求")
	*msg = "收到了您的消息"
	return nil
}

func main() {
	service := micro.NewService(micro.Name("service.greeter"))
	service.Init()
	micro.RegisterHandler(service.Server(), new(Greeter))
	service.Run()
}
