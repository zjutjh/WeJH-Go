package noticeServices

import (
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func CreateRecord(record models.Notice) error {
	result := ndb.Pick().Create(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteNotice(id int) error {
	result := ndb.Pick().Delete(models.Notice{ID: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateNotice(id int, record models.Notice) error {
	result := ndb.Pick().Model(models.Notice{}).
		Select("title", "content").
		Where(&models.LostAndFoundRecord{ID: id}).Updates(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetNoticeBySuperAdmin() ([]models.Notice, error) {
	var record []models.Notice
	result := ndb.Pick().Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetRecordByAdmin(publisher string) ([]models.Notice, error) {
	var record []models.Notice
	result := ndb.Pick().Where(models.Notice{
		Publisher: publisher,
	}).Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetNoticeById(id int) (models.Notice, error) {
	var record models.Notice
	result := ndb.Pick().Where(models.LostAndFoundRecord{
		ID: id,
	}).First(&record)
	if result.Error != nil {
		return models.Notice{}, result.Error
	}
	return record, nil
}
