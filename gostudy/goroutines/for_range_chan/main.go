package main

import (
	"fmt"
	"time"
)

func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	close(c)
}

func main() {
	ch := make(chan int)
	go produce(ch)
	for v := range ch {
		fmt.Println("Recvive: ", v)
	}

}
