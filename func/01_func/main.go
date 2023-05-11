package main

import (
	"fmt"
)

//函数
//函数存在的意义
//函数是一段代码的封装：把一段逻辑抽象出来封装到一个函数中，给他起个名字每次用到的时候直接用函数名调用就可以了
//使用函数能够让代码结构更清晰、更简洁。
//函数的定义
func sum(x int, y int) int {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("没有参数")
}

//没有参数 有返回值
func f3() int {
	//return 0
	ret := 3
	return ret // return后不可省略
}

//返回值可以命名 也可以不命名
//命名的返回值相当于在函数中声明了一个变量
func f4(x, y int) (sum int) {
	sum = x + y
	return sum //使用命名返回值可以return后面省略
}

//多个返回值
func f5() (int, string) {
	return 1, "hello,wrold"
}

//可变参数,a类型是一个切片
func f6(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

//go语言中函数没有默认参数的概念
func main() {
	a := f6(1, 2)
	fmt.Println(a)

}
