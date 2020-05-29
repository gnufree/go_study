package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Printf("连接失败!, err:%v\n", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		sendata, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err)
			break
		}
		sendata = strings.TrimSpace(sendata)
		send := []byte(sendata)
		_, err = socket.Write(send)
		if err != nil {
			fmt.Printf("发送数据失败!, err:%v\n", err)
			return
		}
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("接收数据失败!, err:%v\n", err)
			return
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n", data)
	}
	// 发送数据
	// sendata := []byte("hello server!")
	// _, err = socket.Write(sendata)
	// if err != nil {
	// 	fmt.Printf("发送数据失败!, err:%v\n", err)
	// 	return
	// }
	// // 接收数据
	// data := make([]byte, 4096)
	// read, remoteAddr, err := socket.ReadFromUDP(data)
	// if err != nil {
	// 	fmt.Printf("接收数据失败!, err:%v\n", err)
	// 	return
	// }
	// fmt.Println(read, remoteAddr)
	// fmt.Printf("%s\n", data)
}
