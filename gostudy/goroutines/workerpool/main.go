package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Number int32
	Id     int
}

type Result struct {
	job *Job
	sum int32
}

func calc(job *Job, result chan *Result) {
	var sum int32
	number := job.Number
	for number != 0 {
		tmp := number % 10
		// fmt.Printf("tmp = %d\n", tmp)
		sum += tmp
		// fmt.Printf("sum = %d\n", sum)
		// fmt.Printf("a = %d\n", a)
		number /= 10
		// fmt.Printf("a = %d\n", a)
	}
	// fmt.Printf("sum = %d\n", sum)
	r := &Result{
		job: job,
		sum: sum,
	}

	result <- r
}

func Worker(jobChan chan *Job, resultChan chan *Result) {
	for job := range jobChan {
		calc(job, resultChan)
	}
}

func startWorkerPool(num int, jobChan chan *Job, resultChan chan *Result) {

	for i := 0; i < num; i++ {
		go Worker(jobChan, resultChan)
	}
}

func printResult(resultChan chan *Result) {
	for result := range resultChan {
		fmt.Printf("job id:%v number:%v result:%d\n", result.job.Id, result.job.Number, result.sum)
	}
}
func main() {

	JobChan := make(chan *Job, 1000)
	rusultChan := make(chan *Result, 1000)

	startWorkerPool(128, JobChan, rusultChan)

	go printResult(rusultChan)
	var id int
	for {
		id++
		number := rand.Int31()
		job := &Job{
			Id:     id,
			Number: number,
		}
		JobChan <- job
	}
}
