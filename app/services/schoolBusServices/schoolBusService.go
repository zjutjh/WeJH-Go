package schoolBusServices

import (
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func GetSchoolBusList() ([]models.SchoolBus, error) {
	var bus []models.SchoolBus
	result := ndb.Pick().Find(&bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}

func GetSchoolBus(departure, destination, startTime string, busType models.SchoolBusType) ([]models.SchoolBus, error) {
	var buses []models.SchoolBus
	result := ndb.Pick().Where(models.SchoolBus{
		Departure:   departure,
		Destination: destination,
		StartTime:   startTime,
		Type:        busType,
	}).Find(&buses)
	if result.Error != nil {
		return nil, result.Error
	}
	return buses, nil
}

func CreateSchoolBus(bus models.SchoolBus) error {
	result := ndb.Pick().Create(&bus)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateSchoolBus(id int, bus models.SchoolBus) error {
	result := ndb.Pick().Model(models.SchoolBus{}).Where(
		&models.SchoolBus{
			ID: id,
		}).Updates(&bus)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteSchoolBus(id int) error {
	result := ndb.Pick().Delete(models.SchoolBus{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func RecommendSchoolBus(departure, destination, startTime string, busType models.SchoolBusType) ([]models.SchoolBus, error) {
	var buses []models.SchoolBus
	result := ndb.Pick().Where(
		"departure = ? AND destination = ? AND type = ? AND start_time > ?",
		departure, destination, busType, startTime,
	).Order("start_time desc").First(&buses)
	if result.Error != nil {
		return nil, result.Error
	}
	return buses, nil
}

func GetSchoolBusTimeList(departure, destination string, busType models.SchoolBusType) ([]string, error) {
	var time []string
	result := ndb.Pick().Model(models.SchoolBus{}).Select("DISTINCT start_time").Where(models.SchoolBus{
		Departure:   departure,
		Destination: destination,
		Type:        busType,
	}).Order("start_time").Find(&time)
	if result.Error != nil {
		return nil, result.Error
	}
	return time, nil
}
