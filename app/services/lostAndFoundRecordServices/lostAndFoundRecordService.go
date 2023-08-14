package lostAndFoundRecordServices

import (
	"os"
	"strings"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetAllLostAndFoundRecord(campus, kind string, pageNum, pageSize int) ([]models.LostAndFoundRecord, error) {
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

func GetRecord(campus, kind string, lostOrFound int, pageNum, pageSize int) ([]models.LostAndFoundRecord, error) {
	var record []models.LostAndFoundRecord
	result := database.DB.Where(models.LostAndFoundRecord{
		Campus: campus,
		Kind:   kind,
	}).Where(map[string]interface{}{"type": lostOrFound == 1}).
		Not("is_processed", true).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetAllLostAndFoundTotalPageNum(campus, kind string) (*int64, error) {
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

func GetTotalPageNum(campus, kind string, lostOrFound int) (*int64, error) {
	var pageNum int64
	result := database.DB.Model(models.LostAndFoundRecord{}).Where(models.LostAndFoundRecord{
		Campus: campus,
		Kind:   kind,
	}).Where(map[string]interface{}{"type": lostOrFound == 1}).Not("is_processed", true).Count(&pageNum)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pageNum, nil
}

func GetRecordByAdmin(publisher string, lostOrFound, pageNum, pageSize int) ([]models.LostAndFoundRecord, error) {
	var record []models.LostAndFoundRecord
	result := database.DB.Where(models.LostAndFoundRecord{
		Publisher: publisher,
	}).Where(map[string]interface{}{"type": lostOrFound == 1}).
		Not("is_processed", true).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetTotalPageNumByAdmin(publisher string, lostOrFound int) (*int64, error) {
	var pageNum int64
	result := database.DB.Model(models.LostAndFoundRecord{}).Where(models.LostAndFoundRecord{
		Publisher: publisher,
	}).Where(map[string]interface{}{"type": lostOrFound == 1}).
		Not("is_processed", true).Count(&pageNum)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pageNum, nil
}

func GetRecordBySuperAdmin(lostOrFound int, pageNum, pageSize int) ([]models.LostAndFoundRecord, error) {
	var record []models.LostAndFoundRecord
	result := database.DB.Where(map[string]interface{}{"type": lostOrFound == 1}).
		Not("is_processed", true).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("publish_time desc").Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func GetTotalPageNumBySuperAdmin(lostOrFound int) (*int64, error) {
	var pageNum int64
	result := database.DB.Model(models.LostAndFoundRecord{}).Where(map[string]interface{}{"type": lostOrFound == 1}).
		Not("is_processed", true).Count(&pageNum)
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

func RemoveImg(record models.LostAndFoundRecord, img1 string, img2 string, img3 string) {
	if record.Img1 != "" && record.Img1 != img1 && record.Img1 != img2 && record.Img1 != img3 {
		_ = os.Remove("./img/" + strings.TrimPrefix(record.Img1, config.GetWebpUrlKey()))
	}
	if record.Img2 != "" && record.Img2 != img1 && record.Img2 != img2 && record.Img2 != img3 {
		_ = os.Remove("./img/" + strings.TrimPrefix(record.Img2, config.GetWebpUrlKey()))
	}
	if record.Img3 != "" && record.Img3 != img1 && record.Img3 != img2 && record.Img3 != img3 {
		_ = os.Remove("./img/" + strings.TrimPrefix(record.Img3, config.GetWebpUrlKey()))
	}
}

func UpdateRecord(id int, record models.LostAndFoundRecord) error {
	result := database.DB.Model(models.LostAndFoundRecord{}).Select("*").
		Omit("id", "type", "publish_time", "publisher").
		Where(&models.LostAndFoundRecord{ID: id}).Updates(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
