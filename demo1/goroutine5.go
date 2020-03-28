package main

import (
	"fmt"
	"math/rand"
)

//workpool
//计算一个数字各个位数相加之和
/*
func worker(num int, ch chan int){
	sum := 0
	for num != 0{
		tmp := num % 10
		sum += tmp
		num /= 10
	}
	ch <- sum
}



func main() {

	ch := make(chan int, 1)
	num := 123
	worker(num,ch)
	Resultf := <- ch
	fmt.Println(Resultf)
}

*/

type Job struct {
	Number int
	Id int
}

type Result struct {
	job *Job
	sum int
}

func calc(job *Job, result chan *Result){
	var sum int
	//number := job.Number 在GO中都是值传递，这句会把job.Number数据拷贝到number,下面做运算的时候会
	//只对number操作，而不影响job.Number。如果直接用job.Number最后number /= 10的运算会让job.Number变成0(传过的是个指针,Job.Number会改变)
	//这些会导致下面printResult()函数打印job.Number数据时为0
	number := job.Number
	for number != 0{
		tmp := number % 10
		sum += tmp
		number /= 10
	}
	r := &Result{
		job:    job,
		sum: sum,
	}
	
	result <- r

}


func Worker(jobChan chan *Job, resultChan chan *Result){
	for job := range jobChan {
		calc(job, resultChan)
	}
}

func startWorkPool(num int, jobChan chan *Job, resultChan chan *Result){
	for i := 0; i < num; i++ {
		go Worker(jobChan, resultChan)
	}


}

func printResult(resultChan chan *Result)()  {
	for ersult := range resultChan{

		fmt.Printf("job id:%v,number:%v, result:%d\n", ersult.job.Id, ersult.job.Number, ersult.sum)


	}
}
func main()  {
	jobChan := make(chan *Job, 1000)
	resultChan := make(chan *Result, 1000)

	startWorkPool(128, jobChan, resultChan)

	go printResult(resultChan)
	var id int
	for  {
		id ++
		number := rand.Int()

		job := &Job{
			Number: number,
			Id:     id,
		}
		jobChan <- job
	}
}

