package adminServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetBindStatus(studentID string) (zfStatus, ouathStatus bool, err error) {
	var user models.User
	zfStatus = false
	ouathStatus = false
	result := database.DB.Model(models.User{}).Where("student_id = ?", studentID).First(&user)
	if result.Error != nil {
		return zfStatus, ouathStatus, result.Error
	}
	if user.ZFPassword != "" {
		zfStatus = true
	}
	if user.OauthPassword != "" {
		ouathStatus = true
	}
	return zfStatus, ouathStatus, nil
}
