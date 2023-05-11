package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var _defaultLogger = logrus.New()
var WithFields = _defaultLogger.WithFields
var WithContext = _defaultLogger.WithContext
var Traceln = _defaultLogger.Traceln
var Tracef = _defaultLogger.Tracef
var Debugf = _defaultLogger.Debugf
var Debugln = _defaultLogger.Debugln
var Printf = _defaultLogger.Printf
var Println = _defaultLogger.Println
var Infof = _defaultLogger.Infof
var Infoln = _defaultLogger.Infoln
var Warnf = _defaultLogger.Warnf
var Warnln = _defaultLogger.Warnln
var Errorf = _defaultLogger.Errorf
var Errorln = _defaultLogger.Errorln
var Panicf = _defaultLogger.Panicf
var Paincln = _defaultLogger.Panicln
var Fatalf = _defaultLogger.Fatalf

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339Nano})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.TraceLevel)
	_defaultLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	_defaultLogger.SetOutput(os.Stdout)
	_defaultLogger.SetLevel(logrus.TraceLevel)

}
func main() {
	// log.SetReportCaller(true)
	// log.Println("hello,world")
	var a *log.Entry
	fmt.Printf("a:%#v\n", a)
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.Info("hello world")
	// log.WithFields(log.Fields{
	// 	"omg":    true,
	// 	"number": 122,
	// }).Warn("The group's number increased tremendously!")

	// log.WithFields(log.Fields{
	// 	"omg":    true,
	// 	"number": 100,
	// }).Fatal("The ice breaks!")

	_defaultLogger.Warnf("xxx")
	Errorf("aaaa")

}
