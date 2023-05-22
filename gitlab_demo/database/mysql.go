package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() (err error) {
	fmt.Println("InitMySQL...")

	dsn := "root:Unipal666@tcp(172.168.1.10:3306)/jenkins"
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

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
func AddItem(item Item) (int64, error) {
	i, err := insertItem(item)
	return i, err
}
func insertItem(item Item) (int64, error) {
	sqlStr := `INSERT INTO item (code_id,app_name, app_group, app_type,ssh_url_to_repo, http_url_to_repo) VALUES (?,?,?,?,?,?);`
	return ModifyDB(sqlStr, item.CodeID, item.AppName, item.AppGroup, item.AppType, item.SSHURLToRepo, item.HTTPURLToRepo)
}

func QueryMultiRow(item Item) (err error) {
	sqlStr := "select code_id,app_name,app_group, http_url_to_repo from item ;"
	rows, err := db.Query(sqlStr)
	defer rows.Close() //关闭rows
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("rows:%v\n", rows)
	//循环取值
	for rows.Next() {
		err := rows.Scan(&item.CodeID, &item.AppName, &item.AppGroup, &item.HTTPURLToRepo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v,type:%T\n", item, item)
	}
	return
}

// 查询单行
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// 查询多行
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
func QueryItemWithName() (Item, error) {
	var item Item
	return item, nil
}
func QueryItemWithCon(sql string) ([]Item, error) {
	sql = "select code_id,app_name,app_group, http_url_to_repo from item " + sql
	rows, err := QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.CodeID, &item.AppName, &item.AppGroup, &item.HTTPURLToRepo)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	return items, nil
}
func GetItemFromName() ([]Item, error) {
	return QueryItemWithCon("")
}
func SelectItemByWhere() {

}
