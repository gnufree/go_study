package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" db:"db"`
	Sex   int
	Age   int
	Score float32
	// xxx   int
}

func reflect_exp(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	kind := t.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Printf("s is int64\n")
	case reflect.String:
		fmt.Printf("s is string\n")
	case reflect.Struct:
		fmt.Printf("s is struct\n")
		fmt.Printf("field num of s is %d\n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("name:%s type:%v value:%v\n",
				t.Field(i).Name, field.Type(), field.Interface())
		}
	default:
		fmt.Printf("default\n")
	}
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) Print() {
	fmt.Printf("通过反射调用方法 %v\n", s)
}

func reflect_method(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	v.Elem().Field(0).SetString("st01")
	v.Elem().FieldByName("Sex").SetInt(1)
	v.Elem().Field(2).SetInt(18)
	v.Elem().Field(3).SetFloat(99.9)
	// reflect_exp(s)
	fmt.Printf("s: %#v\n", a)
	fmt.Printf("struct student have %d methods\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("struct %d method, name: %s type: %v\n", i, method.Name, method.Type)
	}

	m1 := v.MethodByName("Print")
	var args []reflect.Value
	m1.Call(args)

	m2 := v.MethodByName("SetName")
	var args2 []reflect.Value
	name := "user01"
	nameVal := reflect.ValueOf(name)
	fmt.Printf("nameVal type: %v value: %v\n", nameVal.Type(), nameVal)
	args2 = append(args2, nameVal)
	m2.Call(args2)

	m1.Call(args)
}
func main() {
	var s Student
	// reflect_method(&s)
	s.SetName("xxx")
	v := reflect.ValueOf(&s)
	t := v.Type()

	field0 := t.Elem().Field(0)
	fmt.Printf("tag json= %s\n", field0.Tag.Get("json"))
	fmt.Printf("tag db= %s\n", field0.Tag.Get("db"))

}
