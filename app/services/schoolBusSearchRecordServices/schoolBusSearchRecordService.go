package schoolBusSearchRecordServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetRecord(username string) (models.SchoolBusSearchRecord, error) {
	var record models.SchoolBusSearchRecord
	result := database.DB.Where(models.SchoolBusSearchRecord{
		Username: username,
	}).First(&record)
	if result.Error != nil {
		return models.SchoolBusSearchRecord{}, result.Error
	}
	return record, nil
}

func CreateRecord(record models.SchoolBusSearchRecord) error {
	result := database.DB.Create(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRecord(id int, record models.SchoolBusSearchRecord) error {
	result := database.DB.Model(models.SchoolBusSearchRecord{}).Where(
		&models.SchoolBusSearchRecord{
			ID: id,
		}).Updates(&record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteRecord(id int) error {
	result := database.DB.Delete(models.SchoolBusSearchRecord{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
