package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//go 日志库
func main() {
	//打开文件
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	//写到文件
	log.SetOutput(fileObj)
	


	for {
		log.Println("这是一条测试日志")
		time.Sleep(time.Second * 3)
	}

}

//日志需求分析
//1、支持往不同的地方输出日志
//2、日志分级别
//Debug、Trace、Info、Warning、Error、Fatal
//3、日志要支持开关控制
//4、完整的日志记录要有时间、行号、文件名、日志级别、日志信息
//5、日志文件切割
