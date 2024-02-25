package suppliesServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func CreateQA(QA models.QA) error {
	result := database.DB.Create(&QA)
	return result.Error
}

func GetQAListByPublisher(publisher string) ([]models.QA, error) {
	var QAList []models.QA
	result := database.DB.Where("publisher = ?", publisher).Find(&QAList)
	if result.Error != nil {
		return nil, result.Error
	}
	return QAList, nil
}

func GetQAListBySuperAdmin() ([]models.QA, error) {
	var QAList []models.QA
	result := database.DB.Order("publisher").Find(&QAList)
	if result.Error != nil {
		return nil, result.Error
	}
	return QAList, nil
}

func GetQAbyID(id int) (*models.QA, error) {
	var QA models.QA
	result := database.DB.Where(&models.QA{
		ID: id,
	}).First(&QA)
	if result.Error != nil {
		return nil, result.Error
	}
	return &QA, nil
}

func UpdateQA(QA models.QA) error {
	result := database.DB.Model(&QA).Updates(QA)
	return result.Error
}

func DeleteQA(id int) error {
	result := database.DB.Delete(models.QA{
		ID: id,
	})
	return result.Error
}
