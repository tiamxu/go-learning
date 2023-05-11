package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

//使用值接收者实现了接口的所有方法
// func (c cat) move() {
// 	fmt.Println("走猫步")
// }
// func (c cat) eat(s string) {
// 	fmt.Printf("%v吃%v\n", c.name, s)
// }
//使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步")
}
func (c *cat) eat(s string) {
	fmt.Printf("%v吃%v\n", c.name, s)
}
func main() {
	var a1 animal
	c1 := cat{name: "汤姆猫", feet: 4}   //cat 值类型
	c2 := &cat{name: "黑猫警长", feet: 2} //*cat 指针类型
	a1 = &c1                          //实现animal这个接口的是cat的指针类型
	fmt.Printf("%v\n", a1)
	a1 = c2
	fmt.Printf("%v\n", a1)

}

//使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
//指针接收者实现接口只能存结构体指针
