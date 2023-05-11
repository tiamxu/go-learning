package main

import "fmt"

//构造函数：返回一个结构体变量的函数
type person struct {
	name string
	age  int
}

//构造函数: 约定成俗用new开头
// 返回的是结构体还是结构体指针，一般看内存占用
//当结构体比较大的时候 尽量使用结构体指针，减少程序的内存开销
func newPerson(x string, y int) person {
	return person{
		name: x,
		age:  y,
	}
}

func main() {
	p := newPerson("小明", 100)
	fmt.Printf("%v\n", p)
}
