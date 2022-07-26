package userServices

import (
	"time"
	"wejh-go/app/models"
	"wejh-go/app/services/funnelServices"
	"wejh-go/config/database"
)

func SetZFPassword(user *models.User, password string) error {
	user.ZFPassword = password
	_, err := funnelServices.GetExam(user, string(rune(time.Now().Year())), "3")
	if err != nil {
		return err
	}
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
	return nil
}

func SetCardPassword(user *models.User, password string) error {
	user.CardPassword = password
	_, err := funnelServices.GetCardBalance(user)
	if err != nil {
		return err
	}
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
	return nil
}

func SetLibraryPassword(user *models.User, password string) error {
	user.LibPassword = password
	_, err := funnelServices.GetCurrentBorrow(user)
	if err != nil {
		return err
	}
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
	return nil
}

func SetPhoneNum(user *models.User, phoneNum string) {
	user.PhoneNum = phoneNum
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
}

func SetYxyUid(user *models.User, yxyUid string) {
	user.YXYUid = yxyUid
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
}
