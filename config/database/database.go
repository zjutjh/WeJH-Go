package database

import (
	"fmt"

	"github.com/zjutjh/mygo/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() { // 初始化数据库
	user := config.Pick().GetString("db.username")
	pass := config.Pick().GetString("db.password")
	port := config.Pick().GetString("db.port")
	host := config.Pick().GetString("db.host")
	name := config.Pick().GetString("db.database")

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

}
