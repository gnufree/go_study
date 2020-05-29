package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	var s string
	s = <-ch
	ch <- "hello"
	ch <- "world"
	ch <- "!"
	s1 := <-ch
	s2 := <-ch
	fmt.Printf("s1 =%s\n", s1)
	fmt.Printf("s2 =%s\n", s2)
	fmt.Printf("s =%s\n", s)
}
