package main

import "fmt"

//自定义类型和类型别名
//type后面跟的是类型
type myInt int     //自定义类型
type yourint = int //类型别名,其他如：rune

func main() {
	// var a myInt
	// a = 100
	var a myInt = 100
	fmt.Printf("%d, %T\n", a, a)
	var n yourint = 200
	fmt.Printf("%d, %T\n", n, n)
	var c rune = '中'
	fmt.Printf("%c, %T\n", c, c)

}
