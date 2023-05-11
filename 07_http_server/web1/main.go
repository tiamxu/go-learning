package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func getData(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "Alice", Age: 28, City: "New York"}

	jsonData, err := json.Marshal(&person)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Printf("data:%v\n", string(jsonData))
	w.Write(jsonData)

}
func main() {
	http.HandleFunc("/", getData)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Printf("http server failed,err:%v\n", err)
		return
	}
}
