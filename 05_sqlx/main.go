package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type City struct {
	Id         int
	Name       string
	Population int
}

func initDb() (err error) {
	dsn := "root:123456@tcp(localhost:3306)/test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)
	return
}
func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init mysql error,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功!")
	//查询单条
	sqlStr := `select * from cities where id = ?;`
	var c City
	err = db.Get(&c, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, error:%v\n", err)
		return
	}
	fmt.Printf("city:%#v\n", &c)
	//插叙多条
	sqlStr2 := `select * from cities where id > ?;`
	var cityList []City
	err = db.Select(&cityList, sqlStr2, 0)
	if err != nil {
		fmt.Printf("select failed, error:%v\n", err)
		return
	}
	fmt.Printf("cityList:%#v\n", &cityList)

}
