package main

import (
	"fmt"
)

func main() {
	// //指针
	// //1、&：取地址
	// n := 10
	// p := &n //p 是一个指针
	// fmt.Printf("%v,%T\n", p, p)
	// //2、*：根据地址取值
	// m := *p
	// fmt.Printf("%v,%T\n", m, m)
	// var a *int
	a := new(int)
	*a = 100
	fmt.Println(*a)

	// var b map[string]int
	var b = make(map[string]int, 10)
	b["徐亮"] = 100
	fmt.Println(b)
	//make 和new用来分配内存
	//new用来返回指针类型
}
