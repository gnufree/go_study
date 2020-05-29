package main

import (
	"fmt"
	"github.com/sony/sonyflake"

)

func main()  {
	setting := sonyflake.Settings{}
	sk := sonyflake.NewSonyflake(setting)

	fmt.Println(sk.NextID())

}



