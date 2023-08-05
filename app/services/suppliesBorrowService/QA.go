package suppliesBorrowService

import (
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/config/database"

	"github.com/gin-gonic/gin"
)

// 获取身份
func GetIdentity(c *gin.Context) *string {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		c.AbortWithError(200, apiException.NotLogin)
		return nil
	}

	var identity string
	if user.Type == models.StudentAffairsCenter {
		identity = "学生事务大厅"
	} else if user.Type == models.Admin {
		identity = "Admin"
	} else {
		c.AbortWithError(200, apiException.ServerError)
		return nil
	}

	return &identity
}

func CreateQA(QA models.QA) error {
	result := database.DB.Create(&QA)
	return result.Error
}

// 根据发布组织获取 QA
func GetQAListByPublisher(publisher string) ([]models.QA, error) {
	var QAList []models.QA
	result := database.DB.Where("publisher = ?", publisher).Find(&QAList)
	if result.Error != nil {
		return nil, result.Error
	}
	return QAList, nil
}

// 获取所有问答
func GetQAListBySuperAdmin() ([]models.QA, error) {
	var QAList []models.QA
	result := database.DB.Order("publisher").Find(&QAList)
	if result.Error != nil {
		return nil, result.Error
	}
	return QAList, nil
}

// 通过 id 获取 QA
func GetQAbyID(id int) (*models.QA, error) {
	var QA models.QA
	result := database.DB.Where("id = ?", id).First(&QA)
	if result.Error != nil {
		return nil, result.Error
	}
	return &QA, nil
}

// 更新 QA
func UpdateQA(QA models.QA) error {
	result := database.DB.Model(&QA).Updates(QA)
	return result.Error
}

// 删除 QA
func DeleteQA(id int) error {
	result := database.DB.Delete(models.QA{ID: id})
	return result.Error
}
