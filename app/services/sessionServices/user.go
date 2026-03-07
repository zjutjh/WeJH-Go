package sessionServices

import (
	"errors"
	"strconv"
	"wejh-go/app/models"
	"wejh-go/app/services/userServices"

	"github.com/gin-gonic/gin"
	"github.com/zjutjh/mygo/session"
)

func SetUserSession(c *gin.Context, user *models.User) error {
	id := strconv.Itoa(user.ID)
	err := session.SetUid(c, id)
	return err
}

func GetUserSession(c *gin.Context) (*models.User, error) {
	id, _ := session.GetUid(c)
	if id == "" {
		return nil, errors.New("")
	}
	idInt, _ := strconv.Atoi(id)
	user, _ := userServices.GetUserID(idInt)
	if user == nil {
		ClearUserSession(c)
		return nil, errors.New("")
	}
	return user, nil
}

func UpdateUserSession(c *gin.Context) (*models.User, error) {
	user, err := GetUserSession(c)
	if err != nil {
		return nil, err
	}
	err = SetUserSession(c, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ClearUserSession(c *gin.Context) {
	_ = session.DeleteUid(c)
	return
}
