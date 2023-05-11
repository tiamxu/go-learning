package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
//1、序列化：把go语言中的结构体变量 --> json格式的字符串
//2、反序列化：把json格式的字符串  --> go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "狄仁杰",
		Age:  100,
	}

	fmt.Println(p1)
	//序列化
	b, _ := json.Marshal(p1)
	fmt.Printf("%s\n", string(b))
	//反序列化
	str := `{"name":"李白","age":900}`
	// var p2 = new(person)
	var p2 person
	err := json.Unmarshal([]byte(str), &p2) //传指针是为了在json.Unmarshal内修改p2的值
	if err != nil {
		fmt.Printf("Unmarshal failed, %v", err)
	}
	fmt.Printf("%T, %v\n", p2, p2)
	fmt.Printf("%v, %v\n", p2.Name, p2.Age)

}
