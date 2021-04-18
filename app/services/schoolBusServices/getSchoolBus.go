package schoolBusServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetSchoolBusList() ([]models.SchoolBus, error) {
	var bus []models.SchoolBus
	result := database.DB.Find(bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}
