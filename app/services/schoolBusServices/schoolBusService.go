package schoolBusServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetSchoolBusList() ([]models.SchoolBus, error) {
	var bus []models.SchoolBus
	result := database.DB.Find(&bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}

func GetSchoolBus(departure, destination, startTime string, busType models.SchoolBusType) ([]models.SchoolBus, error) {
	var buses []models.SchoolBus
	result := database.DB.Where(models.SchoolBus{
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
	result := database.DB.Create(&bus)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateSchoolBus(id int, bus models.SchoolBus) error {
	result := database.DB.Model(models.SchoolBus{}).Where(
		&models.SchoolBus{
			ID: id,
		}).Updates(&bus)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteSchoolBus(id int) error {
	result := database.DB.Delete(models.SchoolBus{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func RecommendSchoolBus(departure, destination, startTime string, busType models.SchoolBusType) ([]models.SchoolBus, error) {
	var buses []models.SchoolBus
	result := database.DB.Where(
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
	result := database.DB.Model(models.SchoolBus{}).Select("DISTINCT start_time").Where(models.SchoolBus{
		Departure:   departure,
		Destination: destination,
		Type:        busType,
	}).Order("start_time").Find(&time)
	if result.Error != nil {
		return nil, result.Error
	}
	return time, nil
}
