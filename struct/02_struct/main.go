package main

import "fmt"

//结构体是值类型

type person struct {
	name   string
	gender string
}

//go语言中，函数传参永远传的都是值拷贝
func f(x person) {
	x.gender = "女" //修改的是副本的gender
	fmt.Printf("func x: %v\n", x.gender)
}
func m(x *person) {
	// (*x).gender = "女" //根据内存地址找到那个原变量，修改的就是原变量
	x.gender = "女" //语法糖，自动根据指针找对应的变量 等价于上面的(*x)

	// fmt.Printf("func x: %v\n", x.gender)
}
func main() {
	var p person
	fmt.Printf("type:%T, value:%v\n", p, p)

	p.name = "蔡徐坤"
	p.gender = "男"
	fmt.Printf("func main: %v\n", p.gender)
	// f(p)
	// fmt.Printf("func main: %v\n", p.gender)
	m(&p)
	fmt.Printf("func main: %v\n", p.gender)
	//new返回一个指针
	//结构体指针1
	var p1 = new(person)
	p1.name = "李白"
	fmt.Printf("type p1:%T, value p1: %v\n", p1, p1)
	fmt.Printf("type p1:%p\n", p1)  //p1保存的值就是一个内存地址
	fmt.Printf("type p1:%p\n", &p1) //p1的内存地址

	var a int = 100
	b := &a
	fmt.Printf("type a:%T, value a:%v,type b:%T, value b :%v\n", a, a, b, b)
	fmt.Printf("a:%p, b:%p \n", &a, b) //b的值
	fmt.Printf("b:%p \n", &b)          //b的内存地址
	//结构体指针2 ：key-value初始化
	var p2 = &person{
		name:   "小米",
		gender: "男",
	}
	fmt.Printf("p3: %#v\n", p2)
	//结构体指针3：使用值列表初始化，值的顺序要和结构体定义时字段顺序一致。
	var p3 = person{
		"小米",
		"男",
	}
	fmt.Printf("p3: %#v\n", p3)
	//结构体占用一块连续的内存空间
	//构造函数：返回一个结构体变量的函数
}
