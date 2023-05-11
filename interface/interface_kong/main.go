package main

import "fmt"

//空接口
//interface：关键字
//interface{}：空接口类型

//空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("%T,%v\n", a, a)
}
func main() {
	//定义一个map
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "小明"
	m1["age"] = 100
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳", "Rap"}
	fmt.Printf("%v\n", m1)

	show("字符串")
	show(100)
	show(true)
	show('b')
}
