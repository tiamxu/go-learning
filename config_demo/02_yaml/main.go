package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string `mapstructure:"env"`
	LogLevel string `mapstructure:"log_level"`
	Srv      `mapstructure:"srv"`
	DB       `mapstructure:"db"`
}
type Srv struct {
	Network        string `mapstructure:"network"`
	ListenAddress  string `mapstructure:"listen_address"`
	WithProxy      bool   `mapstructure:"with_proxy"`
	WithReflection bool   `mapstructure:"with_reflection"`
}
type DB struct {
	Driver          string `mapstructure:"driver"`
	Database        string `mapstructure:"database"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

func main() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	var c Config
	_ = viper.Unmarshal(&c)
	fmt.Printf("%#v\n", c)

}
