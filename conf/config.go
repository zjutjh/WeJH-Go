package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config = viper.New()

func Init() { // 从配置文件中读取配置文件信息
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./conf")
	Config.WatchConfig() // 自动更新配置
	err := Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置读取失败\n %v", err))
	}
}
