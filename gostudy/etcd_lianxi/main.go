package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main()  {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:	[]string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:",err)
		return
	}
	fmt.Println("connect success.")
	defer cli.Close()
}
