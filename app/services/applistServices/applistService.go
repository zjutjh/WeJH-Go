package applistServices

import (
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func GetAppList(count int) ([]models.AppList, error) {
	var applists []models.AppList
	result := ndb.Pick().Limit(count).Find(&applists)
	if result.Error != nil {
		return nil, result.Error
	}
	return applists, nil
}

func CreateApplist(appList models.AppList) error {
	result := ndb.Pick().Create(&appList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateApplist(appList models.AppList) error {
	result := ndb.Pick().Model(models.AppList{}).Where(
		&models.AppList{
			ID: appList.ID,
		}).Updates(&appList)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteApplist(id int64) error {
	result := ndb.Pick().Delete(models.AppList{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAppListPagination(offset, pagesize int) []models.AppList {
	var applists []models.AppList
	ndb.Pick().Offset(offset).Limit(pagesize).Find(&applists)
	return applists
}
