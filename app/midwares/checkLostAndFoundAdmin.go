package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
)

func CheckLostAndFoundAdmin(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	if user.Type != models.ForU && user.Type != models.Admin &&
		user.Type != models.StudentAffairsCenter {
		apiException.AbortWithException(c, apiException.NotAdmin, nil)
		return
	}
	c.Next()
}
