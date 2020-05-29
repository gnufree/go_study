package main

import (
	"fmt"
)

func main() {
	// 计算： 1+2+3+4+5..+n
	var y = 0
	for i := 1; i <= 100; i++ {
		// 1+2+3+4+5
		y += i
	}
	fmt.Println(y)
	var sum int
	var n = 100
	sum = n * (n + 1) / 2
	fmt.Println(sum)
	// 1.确认面试题意思无误
	// 2.想到所有的解决办法，确认时间复杂度，找出最优的解决方案(时间和空间复杂度最低的)
	// 3.写程序
	// 4.测试
	// Fib: 0,1,1,2,3,5,8,13,21,...
	// F(n) = F(n-1)+F(n-2)
	fib(12)
}

func fib(n int) {
	if n < 2 {
		fmt.Println(n)
	}
	fmt.Printf("n = %d\n", n)
	fmt.Println((n - 1) + (n - 2))
}
