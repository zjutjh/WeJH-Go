package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wejh-go/config/config"
)

var DB *gorm.DB

func Init() { // 初始化数据库
	user := config.Config.GetString("database.user")
	pass := config.Config.GetString("database.pass")
	port := config.Config.GetString("database.port")
	host := config.Config.GetString("database.host")
	name := config.Config.GetString("database.name")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束 提升数据库速度
	})

	if err != nil {
		zap.L().Fatal("Database connect failed", zap.Error(err))
	}

	err = autoMigrate(db)
	if err != nil {
		zap.L().Fatal("Database migrate failed", zap.Error(err))
	}
	DB = db
}
