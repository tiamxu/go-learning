package main

import "fmt"

//类型断言
func assign(a interface{}) {
	fmt.Printf("%T,%v\n", a, a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串:", str)
	}

}
func assign2(a interface{}) {
	fmt.Printf("%T,%v\n", a, a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case bool:
		fmt.Println("是一个bool:", t)
	}

}

func main() {
	assign(100)
	assign("name")
	assign2(true)
}
