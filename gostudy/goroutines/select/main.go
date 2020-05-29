package main

import (
	"fmt"
	"time"
)

func server1(ch chan bool) {
	time.Sleep(1 * time.Second)
	ch <- true
}

// func server2(ch chan string) {
// 	// time.Sleep(time.Second * 3)
// 	ch <- "hello wenjun!"
// }

func main() {
	// timeout := make(chan bool, 1)
	// go server1(timeout)
	// ch := make(chan int)
	// select {
	// case <-ch:
	// case <-timeout:
	// 	fmt.Println("超时")
	// }
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
		// fmt.Println(a)
	default:
		fmt.Println("ch full")

	}

}
