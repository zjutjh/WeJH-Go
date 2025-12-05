package userServices

import (
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"

	"github.com/zjutjh/mygo/ndb"
)

func DelAccount(user *models.User, iid string) error {
	if err := userCenterServices.DelAccount(user.Username, iid); err != nil {
		return err
	}
	result := ndb.Pick().Delete(user)
	return result.Error
}
