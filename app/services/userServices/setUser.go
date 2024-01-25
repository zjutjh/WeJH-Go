package userServices

import (
	"time"
	"wejh-go/app/models"
	"wejh-go/app/services/funnelServices"
	"wejh-go/config/database"
)

func SetZFPassword(user *models.User, password string) error {
	user.ZFPassword = password
	_, err := funnelServices.BindPassword(user, string(rune(time.Now().Year())), "3", "ZF")
	if err != nil {
		return err
	}
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
	return nil
}

func SetOauthPassword(user *models.User, password string) error {
	user.OauthPassword = password
	_, err := funnelServices.BindPassword(user, string(rune(time.Now().Year())), "3", "OAUTH")
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
	user.YxyUid = yxyUid
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
}

func SetDeviceID(user *models.User, deviceID string) {
	user.DeviceID = deviceID
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
}

func DelPassword(user *models.User, passwordType string) {
	switch passwordType {
	case "ZF":
		{
			user.ZFPassword = ""
		}
	case "OAUTH":
		{
			user.OauthPassword = ""
		}
	case "Library":
		{
			user.LibPassword = ""
		}
	}
	EncryptUserKeyInfo(user)
	database.DB.Save(user)
}
