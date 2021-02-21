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
	Config.WatchConfig() // 自动将配置读入Config变量
	err := Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置读取失败\n %v", err))
	}
	// TODO: 下次添加往配置文件中自动写入密钥
}
