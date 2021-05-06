package userServices

import (
	"wejh-go/app/config"
	"wejh-go/app/models"
)

func decryptUserKeyInfo(user models.User) {
	config.GetEncryptKey()

}

func encryptUserKeyInfo(user models.User) {
	config.GetEncryptKey()

}
