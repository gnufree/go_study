package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed , err:%v\n", err)
			break
		}

		str := string(buf[:n])
		fmt.Printf("recv from client, data: %v\n", str)
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Printf("listen failed, err:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:", err)
			continue

		}
		go process(conn)
	}
}
