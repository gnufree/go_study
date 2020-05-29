package main

import (
	"fmt"
	"reflect"
)

func reflect_exp(a interface{}) {
	b := reflect.TypeOf(a)
	fmt.Printf("b is type: %v\n", b)
}

func reflect_value(a interface{}) {
	b := reflect.ValueOf(a)
	// fmt.Printf("b is type: %v value: %v\n", b.Kind(), b)

	k := b.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64, store value is:%d\n", b.Int())
	case reflect.Float64:
		fmt.Printf("a is float64, store value is: %f\n", b.Float())
	}
}

func reflect_set_value(a interface{}) {
	b := reflect.ValueOf(a)
	// fmt.Printf("b is type: %v value: %v\n", b.Kind(), b)

	k := b.Kind()
	switch k {
	case reflect.Int64:
		b.SetInt(100)
		fmt.Printf("a is int64, store value is:%d\n", b.Int())
	case reflect.Float64:
		b.SetFloat(4.2)
		fmt.Printf("a is float64, store value is: %f\n", b.Float())
	case reflect.Ptr:
		b.Elem().SetFloat(6.8)
		fmt.Printf("a is float64, store value is: %f\n", b.Pointer())
	}
}

func main() {
	// a := "123"
	// reflect_exp(a)
	// reflect_value(a)
	var c float64 = 3.4
	// reflect_exp(c)
	// reflect_value(c)
	reflect_set_value(&c)
	fmt.Printf("store value is :%v\n", c)

	/*
		var b *int = new(int)
		*b = 100
	*/

}
