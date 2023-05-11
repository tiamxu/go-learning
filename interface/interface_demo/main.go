package main

import "fmt"

//接口示例2
//定义一个car接口类型
//不管是什么结构体，只要有run方法都是car类型
type car interface {
	run()
}
type falali struct {
	name string
}

type baoshijie struct {
	name string
}

func (f falali) run() {
	fmt.Printf("%s时速70\n", f.name)
}

func (b baoshijie) run() {
	fmt.Printf("%s时速700\n", b.name)
}

//driver函数接收一个car类型的变量
func driver(c car) {
	c.run()
}
func main() {
	var f1 = falali{
		name: "法拉利",
	}
	var b1 = baoshijie{
		name: "保时捷",
	}
	driver(f1)
	driver(b1)

}
