package main

import (
	"fmt"
)

func bubble_sort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	// a := 10
	// b := 20
	// c := Add(a, b)
	// fmt.Println("a + b = ", c)
	var i [8]int = [8]int{8, 3, 2, 9, 4, 6, 10, 0}
	j := bubble_sort(i)
	fmt.Println(j)
}
