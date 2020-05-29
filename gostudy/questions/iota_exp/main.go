package main

import "fmt"

const (
	a = iota
	b = iota
)

const (
	name = "name"
	cc = "hh"
	c = iota
	ggg = "gg"
	d = iota
)

const (
	x = iota
	y
	z
	_
	_
	x1
	y1
)
// 位掩码表达式
// 位运算符 & | ^ << >>
type Allergen int
const (
	IgEggs Allergen = 1 << iota
	IgChocolate
	IgNuts
	IgStawberries
	IgShellfish
)
// 定义数量级
type ByteSize float64
const (
	_ = iota 						// ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)  // 1 << (10 * 1)
	MB								// 1 << (10 * 2)
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(name)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(IgEggs) // 000000001
	fmt.Println(IgChocolate) // 00000010
	fmt.Println(IgNuts) // 00000100
	fmt.Println(IgStawberries) // 00001000
	fmt.Println(IgShellfish) // 00010000
	fmt.Println(IgEggs|IgChocolate|IgShellfish)
	var g1 = 60 		// 00111100
	var g2 = 13 		// 00001101
	var g3 = g1 & g2 	// 00001100 按位与 0和0 为0 0 和1 为0 1和0 为0 1和1为1
	fmt.Println(g3)
	var m1 = 60  		// 00111100
	var m2 = 13			// 00001101
	var m3 = m1 | m2	// 00111101  按位或 0和0 为0， 1和0 为1，1和1为1，0和1位1
	fmt.Println(m3)
	var k1 = 60			// 00111100
	var k2 = 13			// 00001101
	var k3 = k1 ^ k2	// 00110001  按位异或 0和0为0，1和0位1，1和1为0，0和1为1
	fmt.Println(k3)
	fmt.Printf("KB: %v\n",KB)
	fmt.Printf("MB: %v\n",MB)
	fmt.Printf("GB: %v\n",GB)
	fmt.Printf("TB: %v\n",TB)
	fmt.Printf("PB: %v\n",PB)
	fmt.Printf("EB: %v\n",EB)
	fmt.Printf("ZB: %v\n",ZB)
	fmt.Printf("YB: %v\n",YB)
	fmt.Println(1 << 10 ) //1000000000               512 256 128 64 32 16 8 4 2 1
}

/*
	结论：
	1、iota 是golang语言的常量计算器，只能在常量表达式中使用
	2、iota在const关键字出现是被重置为0，const每增加一行常量声明将使iota计数一次。（iota可理解为const语句块中的行索引）
	3、使用iota能简化定义，在定义枚举时很有用。
*/
