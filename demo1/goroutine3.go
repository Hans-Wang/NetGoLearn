package main

import "fmt"

//只能向data发送数据的chan
func written(data chan<- int){
	data <- 10
}

//只能从data里读取数据的chan
func reading(data <-chan int){
	num := <- data

	fmt.Println(num)
}

func producer(c chan int)  {
	for i := 0; i < 10; i ++ {
		c <- i
	}
	close(c)
}

func main() {
	//单向chan

	var ch chan int
	ch = make(chan int)
	defer close(ch)
	go written(ch)
	reading(ch)

	c := make(chan int)
	go producer(c)
	for v := range c {
		fmt.Println("receive:", v)
	}

}
