package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wejh-go/conf"
)

var DB *gorm.DB

func Init() { // 初始化数据库
	// 从配置文件中读取数据库信息
	user := conf.Config.GetString("database.user")
	pass := conf.Config.GetString("database.pass")
	port := conf.Config.GetString("database.port")
	host := conf.Config.GetString("database.host")
	name := conf.Config.GetString("database.name")

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		user,
		pass,
		host,
		port,
		name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束 提升数据库速度
	})
	if err != nil {
		panic(fmt.Errorf("数据库连接错误！ \n %v", err))
	}

	// 开始迁移数据
	migrateUsers(DB) // 迁移用户数据
}
