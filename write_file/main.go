package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFileDemo1() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
	}
	
	fileObj.Write([]byte("test!\n"))
	fileObj.WriteString("测试！")
	fileObj.Close()
}

func writeFileBuffIo() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("这是buffio!") //写到缓存中
	wr.Flush()                  //将缓存中的内容写入文件
}

func writeDemo3() {
	str := "这是ioutil"
	err := ioutil.WriteFile("./xxx.txt", []byte(str), 644)
	if err != nil {
		fmt.Printf("write file failed err:%v", err)
		return
	}

}
func main() {
	// writeFileDemo1()
	// writeFileBuffIo()
	writeDemo3()
}
