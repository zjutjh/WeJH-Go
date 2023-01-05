package customizeHomeServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func CreateCustomizeHome(home models.CustomizeHome) error {
	result := database.DB.Create(&home)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCustomizeHome(username string) (*models.CustomizeHome, error) {
	var home models.CustomizeHome
	result := database.DB.Where(models.CustomizeHome{
		Username: username,
	}).First(&home)
	if result.Error != nil {
		return nil, result.Error
	}
	return &home, nil
}

func UpdateCustomizeHome(id int, home models.CustomizeHome) error {
	result := database.DB.Model(models.CustomizeHome{}).Where(
		&models.CustomizeHome{
			ID: id,
		}).Updates(&home)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
