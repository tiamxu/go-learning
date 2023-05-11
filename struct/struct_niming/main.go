package main

import "fmt"

//匿名字段
//字段比较少也比较简单的场景
//不常用
type person struct {
	string
	int
}

func main() {
	p := person{
		string: "李白",
		int:    100,
	}
	fmt.Println(p)
	fmt.Printf("%v, %v\n", p.string, p.int)
}
