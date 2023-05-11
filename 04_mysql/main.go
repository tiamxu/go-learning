package main

//练习mysql初始化、查询、写入
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/test"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	defer db.Close()
	if err != nil { //dsn格式不正确的时候会报错
		log.Fatal(err)
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Println("mysql connect error:", err)
		return
	}
	var version string
	err = db.QueryRow("SELECT VERSION();").Scan(&version)
	if err != nil {
		fmt.Println("query scan error:", err)
		return
	}
	fmt.Println("mysql connect success:", version)
	//查询
	// res, err := db.Query("select * from cities where id = ?", 2)
	// defer res.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for res.Next() {
	// 	var city City
	// 	err := res.Scan(&city.Id, &city.Name, &city.Population)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%v\n", city)
	// }
	//写入
	sql := "insert into cities(name,population) values('Moscow',1250600)"
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println("insert into error:", err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("insert success,the last inserted row id: %d\n", lastId)
}
