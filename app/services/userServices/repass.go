package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/config/database"
)

func ResetPass(user *models.User, iid, password string) error {
	if err := userCenterServices.ResetPass(user.Username, iid, password); err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))
	user.JHPassword = pass
	EncryptUserKeyInfo(user)
	err := database.DB.Model(models.User{}).Where(
		models.User{
			Username: user.Username,
			ID:       user.ID,
		}).Updates(user).Error
	return err
}
