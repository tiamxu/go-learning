package main

import "fmt"

//嵌套结构体
type address struct {
	province string
	city     string
}
type person struct {
	name string
	age  int
	addr address
}
type company struct {
	name    string
	address //匿名嵌套结构体
	//address:address
}

func main() {
	p1 := person{
		name: "宋江",
		age:  100,
		addr: address{province: "湖北", city: "长沙"},
	}
	fmt.Println(p1)
	fmt.Printf("%v,%v,%v,%v\n", p1.name, p1.age, p1.addr.province, p1.addr.city)
	c1 := company{
		name: "京东",
		address: address{
			province: "广东",
			city:     "深圳",
		},
	}
	fmt.Println(c1)
	fmt.Printf("%v,%v,%v\n", c1.name, c1.province, c1.city)
	//先在自己的结构体找这个字段，找不到再去匿名嵌套结构体中找
}
