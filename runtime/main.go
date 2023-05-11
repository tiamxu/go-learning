package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNo
}
func test() {
	funcName, fileName, lineNo := getInfo(0)
	fmt.Println(funcName) //函数名
	fmt.Println(fileName) //文件名
	fmt.Println(lineNo)   //行号
}
func main() {
	// pc, file, lineNo, ok := runtime.Caller(0)
	// if !ok {
	// 	fmt.Printf("runtime.Caller() failed\n")
	// 	return
	// }
	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName) //函数名
	// fileName := path.Base(file)
	// fmt.Println(fileName) //文件名
	// fmt.Println(lineNo)   //行号
	test()
}
