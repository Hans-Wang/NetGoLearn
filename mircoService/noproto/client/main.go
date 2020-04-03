package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func main()  {
	service := micro.NewService()
	service.Init()
	c:=service.Client()
	request := c.NewRequest("service.greeter", "Greeter.Hello", "小明",
		client.WithContentType("application/json"))
	var response string
	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
