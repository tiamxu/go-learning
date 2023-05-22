package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("hello")
}
func main() {
	go f()
	// fmt.Println("你好你好你好你好你好你好你好你好你好你好")
	fmt.Println("你好")

	time.Sleep(time.Second)
}
