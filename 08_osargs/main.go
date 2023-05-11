package main

import (
	"fmt"
)

func main() {
	// if len(os.Args) > 0 {
	// 	fmt.Println(os.Args)
	// 	for index, value := range os.Args {
	// 		fmt.Printf("args[%d]=%v\n", index, value)
	// 	}
	// }
	type person struct {
		name map[string]string
	}
	var (
		defaultMap = map[string]string{"test": "latest"}
		Map        map[string]string
	)
	// Map := make(map[string]string)
	fmt.Println("length", len(Map))

	if Map == nil {
		Map = defaultMap
		fmt.Println("inout", Map)

	}
	fmt.Println("out", Map)
	fmt.Println(defaultMap)

}

//os.Args是一个存储命令行参数的字符串切片，第一个参数是执行文件名本身
