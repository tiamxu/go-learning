package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sendMsg(url, msg string) {
	// json
	contentType := "application/json"
	// data
	sendData := `{
		"msg_type": "text",
		"content": {
			"text": " ` + msg + `"
		}
	 }`
	// request
	result, err := http.Post(url, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

func main() {
	// webhook地址
	url := "https://open.feishu.cn/open-apis/bot/v2/hook/bf8bb912-bc2e-40ad-9533-fcb8068aa621"

	// msg 消息内容
	appname := "hello"
	status := "OK"
	// msg := "appname: " + appname + "status:" + status
	msg := fmt.Sprintf(`name: %s\n status:%s\n`, appname, status)

	sendMsg(url, msg)
}
