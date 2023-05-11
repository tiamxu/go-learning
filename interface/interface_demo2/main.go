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

func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(s string) {
	fmt.Printf("%v吃%v\n", c.name, s)
}

type chicken struct {
	feet int8
}

func main() {
	var a1 animal
	var c1 = cat{
		name: "小花猫",
		feet: 4,
	}
	a1 = c1
	a1.move()
	a1.eat("鱼")
}
//接口：
//接口是一种类型，是一种特殊的类型，他规定了变量有哪些方法
//在编程中会遇到以下场景
//我不关心一个变量是什么类型，我只关心能调用它的什么方法。
//用来给变量、参数、返回值等设置类型
//接口的实现：
//一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量