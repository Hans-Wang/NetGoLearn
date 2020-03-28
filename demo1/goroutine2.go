package main

import (
	"fmt"
)

/*
func hello(){
	fmt.Println("Hello World")
}

func main() {
	go hello()
	//为了防止主进程先退出，先sleep住
	time.Sleep(time.Second)
	fmt.Println("main运行完毕，程序退出!")
}

 */
func hello(c chan bool){
	fmt.Println("Hello World")
	c <- false
}

func main() {
	var exitChan chan bool
	exitChan = make(chan bool)
	go hello(exitChan)
	//使用chan可以替代sleep
	<-exitChan
	//b := <-exitChan
	//if !b {
	//	fmt.Println("hello()执行失败")
	//}else {
	//	fmt.Println("hello()执行成功")
	//}
	fmt.Println("main运行完毕，程序退出!")
}


