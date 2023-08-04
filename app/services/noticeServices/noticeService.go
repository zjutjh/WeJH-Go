package noticeServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func CreateRecord(record models.Notice) error {
	result := database.DB.Create(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteNotice(id int) error {
	result := database.DB.Delete(models.Notice{ID: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateNotice(id int, record models.Notice) error {
	result := database.DB.Model(models.Notice{}).
		Select("title", "content").
		Where(&models.LostAndFoundRecord{ID: id}).Updates(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetNoticeBySuperAdmin() ([]models.Notice, error) {
	var record []models.Notice
	result := database.DB.Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetRecordByAdmin(publisher string) ([]models.Notice, error) {
	var record []models.Notice
	result := database.DB.Where(models.Notice{
		Publisher: publisher,
	}).Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetNoticeById(id int) (models.Notice, error) {
	var record models.Notice
	result := database.DB.Where(models.LostAndFoundRecord{
		ID: id,
	}).First(&record)
	if result.Error != nil {
		return models.Notice{}, result.Error
	}
	return record, nil
}
