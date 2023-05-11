package main

import (
	"fmt"

	"github.com/spf13/viper"
)

//MysqlConfig mysql配置结构体
type MysqlConfig struct {
	Address  string `mapstructure:"address"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// type RedisConfig struct {
// 	Host     string `ini:"host"`
// 	Port     int    `ini:"port"`
// 	Password string `ini:"password"`
// 	Database int    `ini:"database"`
// }
// type Config struct {
// 	MysqlConfig `ini:"mysql"`
// 	RedisConfig `ini:"redis"`
// }

func main() {
	config := viper.New()
	config.AddConfigPath("./config")
	config.SetConfigName("config")
	config.SetConfigType("ini")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件。。。", err)
		} else {
			fmt.Println("配置文件出错...", err)
		}
	}
	host := config.GetString("mysql.address")
	port := config.GetInt("mysql.port")
	fmt.Println("host:", host)
	fmt.Println("port:", port)

}
