package main

import (
	"fmt"
	"time"
)

func write(ch chan string){
	for {
		select {
		case ch <- "Hello":
			fmt.Println("succ")
		default:
			fmt.Printf("Full\n")
		}
		time.Sleep(time.Millisecond*500)
	}
}

func main() {
	output1 := make(chan string, 5)

	go write(output1)
	for r := range output1 {
		fmt.Printf("recv:%s\n", r)
		time.Sleep(time.Second)
	}

}


//空select表示阻塞了。
//select {
//
//}