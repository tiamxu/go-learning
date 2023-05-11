package main

import "fmt"

//结构体 是值类型
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个person类型的变量p
	var p person
	//通过字段赋值
	p.name = "蔡徐坤"
	p.age = 22
	p.gender = "男"
	p.hobby = []string{"唱", "跳", "Rap"}
	fmt.Printf("type:%T, value:%v\n", p, p)
	//访问变量的字段
	fmt.Println(p.name)
	//匿名结构体,多用于一些临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 18
	fmt.Printf("type:%T, value:%v\n", s, s)
}
