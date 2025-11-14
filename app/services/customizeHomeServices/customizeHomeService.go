package customizeHomeServices

import (
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func CreateCustomizeHome(home models.CustomizeHome) error {
	result := ndb.Pick().Create(&home)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCustomizeHome(username string) (*models.CustomizeHome, error) {
	var home models.CustomizeHome
	result := ndb.Pick().Where(models.CustomizeHome{
		Username: username,
	}).First(&home)
	if result.Error != nil {
		return nil, result.Error
	}
	return &home, nil
}

func UpdateCustomizeHome(id int, home models.CustomizeHome) error {
	result := ndb.Pick().Model(models.CustomizeHome{}).Where(
		&models.CustomizeHome{
			ID: id,
		}).Updates(&home)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
