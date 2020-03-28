package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func print(n int, ch chan bool){
	fmt.Println("goroutine start:", n)
	time.Sleep(time.Second * 2)
	fmt.Println("goroutine end:", n)
	ch <- true

}

func main() {
	no := 3
	ch := make(chan bool)

	for i := 0; i< no; i++ {
		go print(i, ch)
	}

	for i := 0; i< no; i++ {
		<- ch
	}
}

*/

//使用waitgroup实现

func print(n int, wg *sync.WaitGroup){
	fmt.Println("goroutine start:", n)
	time.Sleep(time.Second * 2)
	fmt.Println("goroutine end:", n)
	wg.Done()

}

func main() {

	var wg sync.WaitGroup
	no := 3

	for i := 0; i< no; i++ {
		wg.Add(1)
		go print(i, &wg)
	}
	wg.Wait()

	fmt.Println("全部执行完毕。")
}