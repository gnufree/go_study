package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("numbers: %d \n", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("alphabets: %c \n", i)
	}
}

func main() {
	go numbers()
	go alphabets()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
