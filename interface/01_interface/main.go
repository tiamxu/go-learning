package main

import "fmt"

//接口：接口是一个类型
//定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}
type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}
func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}
func da(s speaker) {
	//接受一个参数，传进来什么，我就打什么
	s.speak()
}
func main() {
	var c1 cat
	var d1 dog
	var p1 person
	// c1.speak()
	// d1.speak()
	// p1.speak()
	da(c1)
	da(d1)
	da(p1)
}
