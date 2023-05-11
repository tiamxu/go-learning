package main

import "fmt"

//结构体模拟实现其他语言中的继承
type animal struct {
	name string
}

type dog struct {
	feet   uint8
	animal //animal拥有的方法，dog此时也有了
}

//给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

//给dog实现一个wang的方法
func (d dog) wang() {
	fmt.Printf("%s汪汪汪!\n", d.name)
}
func main() {
	d1 := dog{
		feet:   4,
		animal: animal{name: "小花狗"},
	}
	d1.wang()
	d1.move()
}
