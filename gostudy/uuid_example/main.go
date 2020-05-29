package main

import (
	"fmt"
	 uuid "github.com/satori/go.uuid"
	"reflect"
)

func main() {
	// Creating UUID Version 4
	// panic on error
	id := uuid.NewV4()

	fmt.Printf("%v\n",id)
	fmt.Println("str",id.String())
	fmt.Println("type:", reflect.TypeOf(id.String()))
	fmt.Println("type:", reflect.ValueOf(id))
}