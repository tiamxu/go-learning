package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//通过http请求插入mysql数据
//序列化：结构体-->json格式的字符串
//反序列化: json格式的字符串--> 结构体
//http server:将接收到的json格式字符串，转换为结构体
//结构体-->mysql
type City struct {
	Name       string `json:"name" db:"name" ini:"name"`
	Population int    `json:"population" db:"population" ini:"population"`
}

//打印请求日志
func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed,err:%v\n", err)
		return
	}
	city := string(b) //map字符串
	fmt.Printf("value:%v,type:%T\n", city, city)
	c1 := &City{}
	err = json.Unmarshal([]byte(city), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!", err)
		return
	}
	fmt.Printf("%#v,type:%T,name:%#v,population:%#v\n", c1, c1, c1.Name, c1.Population)
	dsn := "root:123456@tcp(localhost:3306)/test"
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	sql := "insert into cities(name,population) values(?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("prepare insert into error:", err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(c1.Name, c1.Population)
	if err != nil {
		fmt.Println("insert failed,", err)
		return
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("insert success,the last inserted row id: %d\n", lastId)
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}
func main() {
	http.HandleFunc("/post", postHandler)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Printf("http server failed,err:%v\n", err)
		return
	}

}
