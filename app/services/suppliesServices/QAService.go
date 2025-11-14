package suppliesServices

import (
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func CreateQA(QA models.QA) error {
	result := ndb.Pick().Create(&QA)
	return result.Error
}

func GetQAListByPublisher(publisher string) ([]models.QA, error) {
	var QAList []models.QA
	result := ndb.Pick().Where("publisher = ?", publisher).Find(&QAList)
	if result.Error != nil {
		return nil, result.Error
	}
	return QAList, nil
}

func GetQAListBySuperAdmin() ([]models.QA, error) {
	var QAList []models.QA
	result := ndb.Pick().Order("publisher").Find(&QAList)
	if result.Error != nil {
		return nil, result.Error
	}
	return QAList, nil
}

func GetQAbyID(id int) (*models.QA, error) {
	var QA models.QA
	result := ndb.Pick().Where(&models.QA{
		ID: id,
	}).First(&QA)
	if result.Error != nil {
		return nil, result.Error
	}
	return &QA, nil
}

func UpdateQA(QA models.QA) error {
	result := ndb.Pick().Model(&QA).Updates(QA)
	return result.Error
}

func DeleteQA(id int) error {
	result := ndb.Pick().Delete(models.QA{
		ID: id,
	})
	return result.Error
}
