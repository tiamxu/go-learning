package mylogger

//打印日志到终端
import (
	"fmt"
	"time"
)

//往终端写日志相关内容
type ConsoleLogger struct {
	Level LogLevel //日志级别 unit8
}

//NewLog 构造函数,传日志级别
func NewConsoleLogger(levelStr string) ConsoleLogger {
	//将用户传进来的日志级别字符串字符串转为结构体的LogLevel类型
	level := parseLogLevel(levelStr)
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level

}
func (c ConsoleLogger) log(level LogLevel, format string, a ...interface{}) {
	if c.enable(level) {

		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		now := time.Now()
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), parseLogString(level), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	// if DEBUG >= c.Level {
	// 	c.log("DEBUG", format, a...)
	// }
	c.log(DEBUG, format, a...)

}
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)

}
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)

}
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)

}
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	// if ERROR >= c.Level {
	c.log(ERROR, format, a...)
	// }

}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	// if FATAL >= c.Level {
	// 	c.log("FATAL", format, a...)
	// }
	c.log(FATAL, format, a...)

}
