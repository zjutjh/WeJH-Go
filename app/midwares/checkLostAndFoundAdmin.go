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
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.ForU && user.Type != models.Admin &&
		user.Type != models.StudentAffairsCenter {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}
	c.Next()
}
