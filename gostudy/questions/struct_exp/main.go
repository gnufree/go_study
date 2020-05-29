package main

import "fmt"

type People interface {
	Show()
}

type Student struct {
}

func (s *Student) Show() {

}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North","East","South","West"}[d]
}

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo":Math{2,3},
}

func main()  {

	var s *Student
	fmt.Println(s)
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
	/*
	总结：仅当动态值和动态类型都为nil时，接口类型值才为nil，上面的代码，给变量p赋值后，p的动态值是nil，但是动态类型却是*Student，是一个nil指针，所以相等条件不成立。
	 */
	fmt.Println(South)
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
}