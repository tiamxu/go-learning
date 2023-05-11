package mylogger

//定义公共内容
import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

//自定义一个日志库
type LogLevel uint8

//定义一个接口
type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

//声明日志级别常量
const (
	DEBUG LogLevel = iota //依次为0 1 2 3 4 5
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//函数:字符串日志级别 转 整数
func parseLogLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG
	case "trace":
		return TRACE
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return DEBUG

	}
}

//函数parseLogString 整数 转字符串
func parseLogString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"

	}
}

//runtime.Caller()函数 拿到谁调用了我这行代码，拿到函数名 文件名
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //函数名
	fileName = path.Base(file)              //获取文件名
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNo
}
