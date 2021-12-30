package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Pathconfig struct {
	Path string
}

func masin() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("C:\\Users\\jimmy\\GolandProjects\\jimmycmd")
	viper.SetConfigName("test")
	viper.SetConfigType("json")

	viper.ReadInConfig()
	var config Pathconfig
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	fmt.Printf("use config:%s\n", viper.ConfigFileUsed())
	fmt.Println(viper.GetString("path"))
	fmt.Printf("config_path:%s\n", config.Path)
}
