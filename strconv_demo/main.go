package main

import (
	"fmt"
	"strconv"
)

func main() {
	//数字转换成字符串
	i := 96
	str := fmt.Sprintf("%d\n", i)
	fmt.Printf("%T %s\n", str, str)
	s1 := fmt.Sprint("中国北京")
	name := "中国"
	age := 100
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	fmt.Printf("%s,%#v\n", s1, s2)

	s := "1000"
	a, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fmt.Printf("parse failed , err:%v\n", err)
	}
	fmt.Printf("%#v, %T\n", a, a)
	b, _ := strconv.Atoi(s)
	fmt.Printf("%#v %T\n", b, b)
	c := strconv.Itoa(age)
	fmt.Printf("%#v %T\n", c, c)

}
