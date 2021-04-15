package config

import (
	"github.com/spf13/viper"
	"log"
	"wejh-go/exception"
)

var Config = viper.New()

func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.WatchConfig() // 自动将配置读入Config变量
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal(exception.ConfigNotFind, err)
	}
}
