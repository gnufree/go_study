package main

import "fmt"

func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	ch := make(chan int)
	go produce(ch)
	for {
		// v := <-ch

		v, ok := <-ch
		if ok == false {
			fmt.Println("chan is closed")
			break
		}

		fmt.Println("Received: ", v, ok)
	}
}
