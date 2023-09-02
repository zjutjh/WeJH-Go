package userServices

import (
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/utils"
)

func DecryptUserKeyInfo(user *models.User) {
	key := config.GetEncryptKey()
	if user.ZFPassword != "" {
		slt := utils.AesDecrypt(user.ZFPassword, key)
		user.ZFPassword = slt[0 : len(slt)-len(user.JHPassword)]
	}
	if user.LibPassword != "" {
		slt := utils.AesDecrypt(user.LibPassword, key)
		user.LibPassword = slt[0 : len(slt)-len(user.JHPassword)]
	}
	if user.PhoneNum != "" {
		slt := utils.AesDecrypt(user.PhoneNum, key)
		user.PhoneNum = slt[0 : len(slt)-len(user.JHPassword)]
	}
	if user.OauthPassword != "" {
		slt := utils.AesDecrypt(user.OauthPassword, key)
		user.OauthPassword = slt[0 : len(slt)-len(user.JHPassword)]
	}
}

func EncryptUserKeyInfo(user *models.User) {
	key := config.GetEncryptKey()
	if user.ZFPassword != "" {
		user.ZFPassword = utils.AesEncrypt(user.ZFPassword+user.JHPassword, key) //salt
	}
	if user.LibPassword != "" {
		user.LibPassword = utils.AesEncrypt(user.LibPassword+user.JHPassword, key)
	}
	if user.PhoneNum != "" {
		user.PhoneNum = utils.AesEncrypt(user.PhoneNum+user.JHPassword, key)
	}
	if user.OauthPassword != "" {
		user.OauthPassword = utils.AesEncrypt(user.OauthPassword+user.JHPassword, key)
	}
}
