package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type database struct { // 数据库信息
	Host string // 数据库地址
	Name string // 数据库名字
	Port int    // 数据库端口
	User string // 数据库用户名
	Pass string // 数据库密码
}

var DB database

func initDataBase() {
	// 配置文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	// 配置文件读取抛出错误
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件出错了！请检查配置文件！\n %s \n", err))
	}
	// 读取配置文件中的数据
	DB.Host = viper.GetString("database.host")
	DB.Port = viper.GetInt("database.port")
	DB.User = viper.GetString("database.user")
	DB.Pass = viper.GetString("database.pass")
	DB.Name = viper.GetString("database.name")
}
