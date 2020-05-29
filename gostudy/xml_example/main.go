package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Servers []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIp   string `xml:"serverIp"`
}

func main() {
	filename := "./config.xml"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Printf("xml Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
}
