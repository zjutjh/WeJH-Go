package suppliesServices

import (
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/utils"
	"wejh-go/config/database"

	"gorm.io/gorm"
)

func GetPersonalInfoByStudentID(studentID string) (*models.PersonalInfo, error) {
	var info models.PersonalInfo
	result := database.DB.Where(&models.PersonalInfo{
		StudentID: studentID,
	}).First(&info)
	if result.Error == gorm.ErrRecordNotFound {
		info.StudentID = studentID
		return &info, result.Error
	} else if result.Error != nil {
		return nil, result.Error
	}
	aesDecryptInfo(&info)
	return &info, nil
}

func SavePersonalInfo(info models.PersonalInfo) error {
	aesEncryptInfo(&info)
	result := database.DB.Omit("id").Save(&info)
	return result.Error
}

func aesEncryptInfo(info *models.PersonalInfo) {
	key := config.GetEncryptKey()
	info.Contact = utils.AesEncrypt(info.Contact, key)
}

func aesDecryptInfo(info *models.PersonalInfo) {
	key := config.GetEncryptKey()
	info.Contact = utils.AesDecrypt(info.Contact, key)
}
