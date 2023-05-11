package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

var db *sql.DB

// 配置文件结构体
type Config struct {
	Env      string `yaml:"env"`
	LogLevel string `yaml:"log_level"`
	// Srv      `yaml:"srv"`
	DB `yaml:"db"`
}
type Srv struct {
	Network        string `yaml:"network"`
	ListenAddress  string `yaml:"listen_address"`
	WithProxy      bool   `yaml:"with_proxy"`
	WithReflection bool   `yaml:"with_reflection"`
}

type DB struct {
	Driver          string `yaml:"driver"`
	Database        string `yaml:"database"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

// 数据库结构体
type City struct {
	Id         int
	Name       string
	Population int
}

// 方法
func (cfg *Config) Source() string {
	switch strings.ToLower(cfg.DB.Driver) {
	case "mysql":
		return cfg.mysqlSource()
	case "postgres":
		return cfg.postgresSource()
	default:
		return ""

	}
}

func (cfg *Config) mysqlSource() string {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
	return dbSource
}
func (cfg *Config) postgresSource() string {
	dbSource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
	return dbSource
}

// 构造函数
// func newConfig(driverName, username, password string, port, MaxOpenConns, ConnMaxLifetime, MaxIdleConns int) *DB {
// 	return &DB{
// 		Driver:          driverName,
// 		Username:        username,
// 		Password:        password,
// 		Port:            port,
// 		MaxOpenConns:    MaxOpenConns,
// 		MaxIdleConns:    MaxIdleConns,
// 		ConnMaxLifetime: ConnMaxLifetime,
// 	}
// }

// 连接数据库
func Connect(dbConfig *Config) (err error) {
	db, err = sql.Open(dbConfig.DB.Driver, dbConfig.Source())
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(dbConfig.DB.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Second)
	return
}
func initDB(driver, dsn string, moc, mic, cml int) (err error) {
	db, err = sql.Open(driver, dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(moc) //设置数据库连接池的最大连接数
	db.SetMaxIdleConns(mic) //设置空闲连接池中的最大连接数
	// db.SetConnMaxIdleTime(3600)
	db.SetConnMaxLifetime(time.Duration(cml) * time.Second)
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
func main() {
	//声明结构体变量
	var c Config
	filename := "./config/config.yaml"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("yaml文件内容:\n%v\n", string(data))
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
	// fmt.Println(c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
	// fmt.Println(dsn)
	// err = initDB(c.DB.Driver, dsn, c.DB.MaxOpenConns, c.DB.MaxIdleConns, c.DB.ConnMaxLifetime)
	err = Connect(&c)
	if err != nil {
		fmt.Printf("数据库初始化错误,error:%s", err)
	}
	fmt.Println("数据库连接成功")
	queryOneRow(1)

}
