package main

import (
	"fmt"
)

//函数中查找变量的顺序1、现在函数内部查找 2、找不到就函数外面找，一直到全局
//匿名函数
var f1 = func(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(f1(1, 2))
	//匿名函数
	//函数内部没有办法声明带名字的函数
	f2 := func(x, y int) {
		fmt.Println(x + y)
	}
	f2(10, 20)
	//立即执行函数
	//如果只是调用一次的函数，可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)
}

//变量作用域
//1、全局作用域
//2、函数作用域
//2、语句块作用域
