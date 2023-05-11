package main

import (
	"fmt"
)

//闭包
//闭包是一个函数，这个函数包含他外部作用域的变量

//底层原理：
//1、函数可以作为返回值
//2、函数内部查找变量的顺序，先在自己内部找，找不到往外层找
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder(100)
	ret2 := ret(100)
	fmt.Println(ret2)
}
