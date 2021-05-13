package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiExpection"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
)

func CheckAdmin(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	if user.Type != models.Admin {
		_ = c.AbortWithError(200, apiExpection.NotAdmin)
		return
	}
	c.Next()

}
