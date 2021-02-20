package database

import (
	"fmt"
	"gorm.io/gorm"
)

func migrateUsers(db *gorm.DB) {
	// 开始生成各用户信息数据表
	if !db.Migrator().HasTable(&User{}) {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic(fmt.Errorf("数据库迁移失败！\n %v", err))
		}
	}
}
