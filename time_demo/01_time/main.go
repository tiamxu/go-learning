package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// var t int = 1
	fmt.Println(1 * time.Second)
	fmt.Println(time.Duration(1))

}
