package applistServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetAppList(count int) ([]models.AppList, error) {
	var applists []models.AppList
	result := database.DB.Limit(count).Find(&applists)
	if result.Error != nil {
		return nil, result.Error
	}
	return applists, nil
}

func GetAppListPagination(offset, pagesize int) []models.AppList {
	var applists []models.AppList
	database.DB.Offset(offset).Limit(pagesize).Find(&applists)
	return applists
}
