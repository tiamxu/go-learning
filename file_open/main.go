package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromfile1() {
	//打开文件
	fileobj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//记得关闭文件
	defer fileobj.Close()
	tmp := make([]byte, 128)
	for {
		n, err := fileobj.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err:%v\n", err)
			return
		}

		fmt.Printf("读取了%d字节数据", n)
		fmt.Println(string(tmp[:n]))
		// if n < 128 {
		// 	return
		// }
	}
}

func readFromFileBufio() {
	//打开文件
	fileobj, err := os.Open("unipal_device_public_device.csv")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//记得关闭文件
	defer fileobj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileobj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v", err)
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("unipal_device_public_device.csv")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
	}
	fmt.Println(string(ret))
}
func main() {
	// readFromfile1()
	readFromFileBufio()
	// readFromFileByIoutil()
}
