package database

import (
	"gorm.io/gorm"
	"wejh-go/app/models"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Announcement{},
		&models.AppList{},
		&models.SchoolBus{},
		&models.Config{},
		&models.SchoolBusSearchRecord{},
		&models.Lesson{},
		&models.CustomizeHome{})
}
