package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	fmt.Println("ch cap: ", cap(ch))
	fmt.Println("ch len: ", len(ch))
	fmt.Println("recvied ", <-ch)
	fmt.Println("ch len: ", len(ch))
}
