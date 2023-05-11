package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigExample struct {
	CompilerOptions struct {
		Module string `json:"module"`
		Target string `json:"target"`
	} `json:"compilerOptions"`
	Exclude []string `json:"exclude"`
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	var config ConfigExample

	viper.Unmarshal(&config)

	fmt.Printf("%v\n", config.Exclude)
}
