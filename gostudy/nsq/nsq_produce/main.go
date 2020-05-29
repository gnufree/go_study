package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nsqio/go-nsq"
)

var produce *nsq.Producer

func main() {
	nsqAddress := "127.0.0.1:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init produce failed, err:%v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string failed, err:%v\n", err)
			continue
		}
		if data == "stop" {
			break
		}

		err = produce.Publish("order_quere", []byte(data))
		if err != nil {
			fmt.Printf("publish message failed, err:%v\n", err)
			continue
		}
		fmt.Printf("pushlist data:%s succ\n", data)
	}

}

func initProducer(str string) error {
	var err error
	config := nsq.NewConfig()
	produce, err = nsq.NewProducer(str, config)

	if err != nil {
		return err
	}
	return nil
}
