package database

import (
	"wejh-go/app/models"

	"gorm.io/gorm"
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
		&models.CustomizeHome{},
		&models.LostAndFoundRecord{},
		&models.LostKind{},
		&models.Notice{},
		&models.LowBatteryQueryRecord{},
		&models.QA{},
		&models.PersonalInfo{},
		&models.Supplies{},
		&models.BorrowRecord{},
	)
}
