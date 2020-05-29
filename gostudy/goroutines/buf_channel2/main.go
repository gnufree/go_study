package main

import (
	"fmt"
	"time"
)

func write(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("successfully write ", i, "to ch")
	}
	close(c)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value ", v, "from ch")
		time.Sleep(time.Second * 2)
	}
}
