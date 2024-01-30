package userServices

import (
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/config/database"
)

func DelAccount(user *models.User, iid string) error {

	if err := userCenterServices.DelAccount(user.Username, iid); err != nil {
		return err
	}

	result := database.DB.Delete(user)
	return result.Error
}
