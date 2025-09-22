package midwares

import (
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"

	"github.com/gin-gonic/gin"
)

func CheckSuppliesAdmin(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	if user.Type != models.StudentAffairsCenter && user.Type != models.Admin {
		apiException.AbortWithException(c, apiException.NotAdmin, nil)
		return
	}
	c.Next()
}
