package main

import "source/a/mylogger"

var log mylogger.Logger

//测试自己写的日志库
func main() {
	log = mylogger.NewConsoleLogger("")
	// log := mylogger.NewFileLogger("debug", "./", "xx.log", 1024)
	err := "test"
	log.Debug("这是一个debug日志, err:%v", err)
	log.Trace("这是一个trace日志")
	log.Info("这是一个info日志")
	log.Warning("这是一个warning日志")
	log.Error("这是一个error日志")
	log.Fatal("这是一个fatal日志")

}
