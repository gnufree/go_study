package main

import (
	"github.com/gnufree/gostudy/logger"
)

// var log logger.LogInterface

func initLogger(name, logPath, logName string, level string) (err error) {
	m := make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = logName
	m["log_level"] = level
	m["log_split_type"] = "size"
	// 1 初始化Logger
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}

	logger.Debug("init logger success")
	return
	// log = logger.NewFileLogger(level, logPath, logName)
	// log = logger.NewConsoleLogger(level)
	// log.Debug("init logger success")
}

func Run() {
	for {
		logger.Debug("user server is running: /Users/zhiwenjun/workspace/code/golang/jikeshijian/src/github.com/gnufree/gostudy/user_server")
		// time.Sleep(time.Second)
	}
}

func main() {
	/*
		file := log.NewFileLog("./a.log")
		file.LogDebug("This is a debug log")
		file.LogWarn("This is a warn log")
	*/
	/*
			console := log.NewConsoleLog("xxxx")
			console.LogConsoleDebug("This is a debug log")
			console.LogConsoleWarn("This is a warn log")

		log := log.NewFileLog("./a.log")
		log := log.NewConsoleLog("xxx")
		log.LogDebug("This is a debug log")
		log.LogWarn("This is a warn log")
	*/
	initLogger("file", "./", "user_server", "debug")
	Run()
	return

}
