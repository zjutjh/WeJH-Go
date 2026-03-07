package schoolBusSearchRecordServices

import (
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func GetRecord(username string) (models.SchoolBusSearchRecord, error) {
	var record models.SchoolBusSearchRecord
	result := ndb.Pick().Where(models.SchoolBusSearchRecord{
		Username: username,
	}).First(&record)
	if result.Error != nil {
		return models.SchoolBusSearchRecord{}, result.Error
	}
	return record, nil
}

func CreateRecord(record models.SchoolBusSearchRecord) error {
	result := ndb.Pick().Create(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRecord(id int, record models.SchoolBusSearchRecord) error {
	result := ndb.Pick().Model(models.SchoolBusSearchRecord{}).Where(
		&models.SchoolBusSearchRecord{
			ID: id,
		}).Updates(&record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteRecord(id int) error {
	result := ndb.Pick().Delete(models.SchoolBusSearchRecord{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
