package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// type City struct {
// 	Id         int
// 	Name       string
// 	Population int
// }

//'Moscow',1250600
func main() {

	url := "http://127.0.0.1:9091/post"
	contentType := "application/json"
	data := `{"name":"Beijing","population":100001}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(b))

}

//client 和server 交互流程
// client 发送数据
//server 接收数据处理
//server 返回数据给client
//client 处理返回数据
