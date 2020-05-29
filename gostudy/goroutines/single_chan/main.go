package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
	// <-sendch
}

func readdData(sendch <-chan int) {
	data := <-sendch
	fmt.Println(data)
}

func main() {
	chan1 := make(chan int)
	go sendData(chan1)
	readdData(chan1)
}
