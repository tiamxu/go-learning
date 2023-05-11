package main

import (
	"fmt"
)

//defer 延迟处理，后进先出
//go语言中return不是原子操作，分2步执行
//1、返回值复制
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//f2=6
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

//f3=5
func f3() (y int) {
	x := 5
	defer func() {
		x++ //修改的是x
	}()
	return x //返回值 = y = x = 5
}

//f4=6
func f4() (x int) {
	defer func(x int) {
		x++ //函数传参，改变的是函数中的副本
	}(x)
	return 5 //返回值 = x = 5
}
func main() {

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
