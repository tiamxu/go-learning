package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //init()
	// func init() {
	// 	sql.Register("mysql", &MySQLDriver{})
	// }
)

//定义一个全局的DB连接池类型
var db *sql.DB

type City struct {
	Id         int
	Name       string
	Population int
}

func initDB() (err error) {
	dsn := "root:123456@tcp(localhost:3306)/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置空闲连接池中的最大连接数
	db.SetConnMaxIdleTime(3600)
	db.SetConnMaxLifetime(3600)
	return

}
func queryVersion() (err error) {
	var version string
	err = db.QueryRow("SELECT VERSION();").Scan(&version)
	if err != nil {
		return
	}
	fmt.Println("mysql version is:", version)

	return
}

func queryOneRow(id int) (err error) {
	var city City
	//查询单条记录的sql语句
	sqlStr := `select * from cities where id = ?;`
	//执行
	// row := db.QueryRow(sqlStr) //从连接池里拿一个连接出来去数据库查询单条记录
	// //拿到结果
	// row.Scan(&city.Id, &city.Name, &city.Population) //必须对row对象调用scan方法，因为该方法会释放数据库连接
	db.QueryRow(sqlStr, id).Scan(&city.Id, &city.Name, &city.Population)
	//打印结果
	fmt.Printf("%#v,type:%T", city, city)
	return
}
func queryMultiRow(id int) (err error) {
	sqlStr := "select * from cities where id > ?;"
	rows, err := db.Query(sqlStr, id)
	defer rows.Close() //关闭rows
	if err != nil {
		log.Fatal(err)
	}
	//循环取值
	for rows.Next() {
		var city City
		err := rows.Scan(&city.Id, &city.Name, &city.Population)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v,type:%T\n", city, city)

	}
	return
}

//插入
func inserRow(name string, population int) {
	// sqlStr := `insert into cities (name,population) value("深圳",90000)`
	sqlStr := `insert into cities (name,population) value(?,?)`

	ret, err := db.Exec(sqlStr, name, population)
	if err != nil {
		fmt.Printf("insert falied,err:%v\n", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id falied,err:%v\n", err)
		return
	}
	fmt.Printf("insert id=%d\n", id)
}
func updateRow(name string, id int) {
	sqlStr := `update cities set name=? where id=? `
	ret, err := db.Exec(sqlStr, name, id)
	if err != nil {
		fmt.Printf("update falied,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected falied,err:%v\n", err)
		return
	}
	fmt.Printf("update success,affected rows:%d\n", n)

}
func deleteRow(id int) {
	sqlStr := `delete from  cities where id=?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete falied,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete RowsAffected falied,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,affected rows:%d\n", n)

}

//预处理插入多条数据
func prepareInsert(m map[string]int) {
	sqlStr := `insert into cities (name,population) value(?,?)`
	stmt, err := db.Prepare(sqlStr) //把sql语言先发给mysql预处理一下
	if err != nil {
		fmt.Printf("pereare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	//后续只需要拿到stmt去执行操作
	for k, v := range m {
		ret, err := stmt.Exec(k, v) //后续只需要传值
		if err != nil {
			fmt.Printf("preare insert failed, err:%v\n", err)
			return
		}
		id, err := ret.LastInsertId()
		if err != nil {
			fmt.Printf("preare get id  failed, err:%v\n", err)
			return
		}
		fmt.Printf("insert id=%d\n", id)
	}

}
func transaction() {
	//事务执行多个sql
	sqlStr1 := `update cities set population=? where id=?`
	sqlStr2 := `update citie set population=? where id=?`
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		fmt.Printf("begin tx failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec(sqlStr1, 1100, 1)
	if err != nil {
		//回滚
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec(sqlStr2, 1101, 2)
	if err != nil {
		//回滚
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit error,err:%v\n", err)
		return
	}
	fmt.Println("事务执行成功.")

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("mysql init DB error:", err)
		return
	}
	fmt.Println("连接数据库成功!")
	err = queryVersion()
	if err != nil {
		fmt.Println("query error:", err)
	}
	//queryOneRow(1)
	// queryMultiRow(2)
	// inserRow("Shenzheng", 9000)
	// queryMultiRow(2)
	//更新
	// updateRow("深圳", 11)
	//删除
	// deleteRow(8)
	//预处理插入
	// var m = map[string]int{
	// 	"成都": 11000,
	// }
	// prepareInsert(m)
	transaction()

	queryMultiRow(0)

}
