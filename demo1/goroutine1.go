package main

import (
	"fmt"
	"runtime"
	"time"
)

func numPrint(){
	for i := 1; i <= 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(i)
	}
}

func letterPrint(){
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

func main() {
	//计算机的核数和设置具体使用几核
	n := runtime.NumCPU()
	fmt.Println("CPU的核数",n)
	runtime.GOMAXPROCS(1)
	go numPrint()
	go letterPrint()

	time.Sleep( 3* time.Second)
	fmt.Println("main运行完毕，程序退出!")
}