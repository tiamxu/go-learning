package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 发送请求到 web2.com
	response, err := http.Get("http://localhost:9091/")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// 解析响应并获取数据
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 在这里可以处理解析出来的数据
	fmt.Println(string(data))
}
