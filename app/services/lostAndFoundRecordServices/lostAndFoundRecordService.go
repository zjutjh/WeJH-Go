package lostAndFoundRecordServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetAllKindRecord(campus string, pageNum, pageSize int) ([]models.LostAndFoundRecord, error) {
	var record []models.LostAndFoundRecord
	result := database.DB.Where(models.LostAndFoundRecord{
		Campus: campus,
	}).Not("is_processed", true).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetRecord(campus, kind string, pageNum, pageSize int) ([]models.LostAndFoundRecord, error) {
	var record []models.LostAndFoundRecord
	result := database.DB.Where(models.LostAndFoundRecord{
		Campus: campus,
		Kind:   kind,
	}).Not("is_processed", true).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetRecordTotalPageNum(campus, kind string) (*int64, error) {
	var pageNum int64
	result := database.DB.Model(models.LostAndFoundRecord{}).Where(models.LostAndFoundRecord{
		Campus: campus,
		Kind:   kind,
	}).Not("is_processed", true).Count(&pageNum)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pageNum, nil
}

func GetRecordAllKindTotalPageNum(campus string) (*int64, error) {
	var pageNum int64
	result := database.DB.Model(models.LostAndFoundRecord{}).Where(models.LostAndFoundRecord{
		Campus: campus,
	}).Not("is_processed", true).Count(&pageNum)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pageNum, nil
}

func GetRecordByAdmin(publisher string, pageNum, pageSize int) ([]models.LostAndFoundRecord, error) {
	var record []models.LostAndFoundRecord
	result := database.DB.Where(models.LostAndFoundRecord{
		Publisher: publisher,
	}).Not("is_processed", true).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetRecordTotalPageNumByAdmin(publisher string) (*int64, error) {
	var pageNum int64
	result := database.DB.Model(models.LostAndFoundRecord{}).Where(models.LostAndFoundRecord{
		Publisher: publisher,
	}).Not("is_processed", true).Count(&pageNum)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pageNum, nil
}

func GetKindList() ([]models.LostKind, error) {
	var kinds []models.LostKind
	result := database.DB.Where(models.LostKind{}).Find(&kinds)
	if result.Error != nil {
		return nil, result.Error
	}
	return kinds, nil
}

func GetRecordById(id int) (models.LostAndFoundRecord, error) {
	var record models.LostAndFoundRecord
	result := database.DB.Where(models.LostAndFoundRecord{
		ID: id,
	}).First(&record)
	if result.Error != nil {
		return models.LostAndFoundRecord{}, result.Error
	}
	return record, nil
}

func CreateRecord(record models.LostAndFoundRecord) error {
	result := database.DB.Create(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRecord(id int, record models.LostAndFoundRecord) error {
	result := database.DB.Model(models.LostAndFoundRecord{}).
		Select("type", "campus", "kind", "content", "is_processed").
		Where(&models.LostAndFoundRecord{ID: id}).Updates(&record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
