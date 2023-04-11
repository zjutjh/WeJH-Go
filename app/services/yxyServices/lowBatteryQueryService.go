package yxyServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func InsertRecord(record models.LowBatteryQueryRecord) error {
	result := database.DB.Create(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
