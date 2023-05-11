package main

import "fmt"

//接口嵌套
type animal interface {
	mover
	eater
}

//同一个结构体可以实现多个接口
type mover interface {
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int8
}

//cat实现了mover接口
func (c *cat) move() {
	fmt.Printf("%v走猫步。\n", c.name)
}

//cat实现了eater接口
func (c *cat) eat(s string) {
	fmt.Printf("%v吃%s\n", c.name, s)
}
func main() {
	var a1 animal
	var m1 mover
	var e1 eater
	c1 := &cat{name: "小丁猫", feet: 4}
	a1 = c1
	m1 = c1
	e1 = c1
	fmt.Printf("%v\n", c1)
	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", e1)
	a1.eat("鱼")
	a1.move()
	m1.move()
	e1.eat("猫粮")


}
