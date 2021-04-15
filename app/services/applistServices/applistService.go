package applistServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetAppList(count int) []models.AppList {
	var applists []models.AppList
	database.DB.Limit(count).Find(&applists)
	return applists
}

func GetAppListPagination(offset, pagesize int) []models.AppList {
	var applists []models.AppList
	database.DB.Offset(offset).Limit(pagesize).Find(&applists)
	return applists
}
