package database

import (
	"fmt"
	"gorm.io/gorm"
	"wejh-go/app/models"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Announcement{}, &models.AppList{}, &models.SchoolBus{})
}
func migrateUsers(db *gorm.DB) {
	// 开始生成各用户信息数据表
	if !db.Migrator().HasTable(&models.User{}) {
		err := db.Migrator().CreateTable(&models.User{})
		if err != nil {
			panic(fmt.Errorf("数据库迁移失败！\n %v", err))
		}
	}
}
func migrateAppList(db *gorm.DB) {
	// 开始生成各用户信息数据表
	if !db.Migrator().HasTable(&models.AppList{}) {
		err := db.Migrator().CreateTable(&models.AppList{})
		if err != nil {
			panic(fmt.Errorf("数据库迁移失败！\n %v", err))
		}
	}
}
