package suppliesServices

import (
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/utils"

	"github.com/zjutjh/mygo/ndb"
	"gorm.io/gorm"
)

func GetPersonalInfoByStudentID(studentID string) (*models.PersonalInfo, error) {
	var info models.PersonalInfo
	result := ndb.Pick().Where(&models.PersonalInfo{
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
	result := ndb.Pick().Omit("id").Save(&info)
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
