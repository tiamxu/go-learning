package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string //日志保存路径
	fileName    string //日志文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 //日志大小
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxSize int64) *FileLogger {
	level := parseLogLevel(levelStr)

	f1 := &FileLogger{
		Level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxSize,
	}
	err := f1.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return f1

}

//初始化文件方法
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

//判断是否需要记录该日志
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level

}

//判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v", err)
		return false
	}
	//如果当前文件大小 大于等于日志文件最大值 就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

//切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割文件
	nowStr := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get split file info failed,err:%v", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) //拿到当前的日志完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	//1、关闭当前文件
	file.Close()
	//2、备份一下 rename
	os.Rename(logName, newLogName)
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	//4、将打开的新日志文件对象赋值给f.fileObj对象
	return fileObj, nil
}

//记录日志的方法
func (f *FileLogger) log(level LogLevel, format string, a ...interface{}) {
	if f.enable(level) {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		now := time.Now()
		if f.checkSize(f.fileObj) {
			//需要切割文件
			//1、关闭当前文件
			// f.fileObj.Close()
			// //2、备份一下 rename
			// nowStr := time.Now().Format("20060102150405")
			// logName := path.Join(f.filePath, f.fileName) //拿到当前的日志完整路径
			// newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
			// os.Rename(logName, newLogName)
			// //3.打开一个新的日志文件
			// fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
			// if err != nil {
			// 	fmt.Printf("open new log file failed, err:%v\n", err)
			// 	return
			// }
			// //4、将打开的新日志文件对象赋值给f.fileObj对象
			// f.fileObj = fileObj
			newFile, err := f.splitFile(f.fileObj)

			if err != nil {
				fmt.Printf("split log file error , err:%v", err)
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), parseLogString(level), fileName, funcName, lineNo, msg)
		if level > ERROR {
			if f.checkSize(f.errFileObj) {
				fileObj, err := f.splitFile(f.errFileObj)

				if err != nil {
					fmt.Printf("split error log file error , err:%v", err)
				}
				f.errFileObj = fileObj
			}
			//如果日志级别大于等于ERROR级别，还要在err日志文件记录
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), parseLogString(level), fileName, funcName, lineNo, msg)

		}

	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)

}
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)

}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
