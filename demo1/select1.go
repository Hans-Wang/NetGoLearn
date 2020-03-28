package main

import (
	"fmt"
	"time"
)

func server1(ch chan string)  {
	time.Sleep(time.Second * 6)
	ch <- "server1"
}

func server2(ch chan string)  {
	time.Sleep(time.Second * 3)
	ch <- "server2"
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)
	/*
	s1 := <- ch1
	fmt.Println("server1:", s1)

	s2 := <- ch2
	fmt.Println("server2:", s2)
	 */


	select {
	case s1 := <- ch1:
		fmt.Println("server1:", s1)
	case s2 := <- ch2:
		fmt.Println("server2:", s2)
	}
}
