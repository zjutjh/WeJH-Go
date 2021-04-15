package sessionServices

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"wejh-go/app/models"
	"wejh-go/app/services/userServices"
)

func SetUserSession(c *gin.Context, user *models.User) error {
	webSession := sessions.Default(c)
	webSession.Set("id", user.ID)
	return webSession.Save()
}

func GetUserSession(c *gin.Context) (*models.User, error) {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return nil, errors.New("")
	}
	user, _ := userServices.GetUserID(id.(int))

	return user, nil
}
