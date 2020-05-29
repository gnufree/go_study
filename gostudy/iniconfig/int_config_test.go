package iniconfig

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port uint   `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

func TestIniConfig(t *testing.T) {
	fmt.Println("hello")
	// t.Errorf("hello")
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("read file failed")
	}

	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("unmarshal failed, err: %v", err)
	}

	confData, err := Marshal(conf)
	if err != nil {
		t.Errorf("marshal failed, err:%v", err)
	}
	t.Logf("marshal success, conf:%s", string(confData))

	t.Logf("unmarshal success, conf:%#v, port:%v", conf, conf.ServerConf.Port)

	MarshalFile("./a.ini", conf)

}

func TestIniConfigFile(t *testing.T) {
	filename := "./test.conf"
	var conf Config
	conf.ServerConf.Ip = "localhost"
	conf.ServerConf.Port = 88888
	err := MarshalFile(filename, conf)
	if err != nil {
		t.Errorf("marshal failed, err: %v", err)
	}

	var conf2 Config
	err = UnMarshalFile(filename, &conf2)
	if err != nil {
		t.Errorf("unmarshal failed, err: %v", err)
	}
	t.Logf("unmarshal success, conf:%#v", conf2)

}
