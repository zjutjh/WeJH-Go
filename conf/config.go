package conf

import "github.com/spf13/viper"

type Database struct { // 数据库信息
	host string // 数据库地址
	name string // 数据库名字
	port int    // 数据库端口
	user string // 数据库用户名
	pass string // 数据库密码
}

func init() { // 从配置文件中读取配置文件信息
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 从当前目录读取
}
