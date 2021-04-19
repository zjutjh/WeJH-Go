package midware

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
)

func CheckAdmin(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.NotLogin, nil)
		c.Abort()
		return
	}
	if user.Type != models.Admin {
		utils.JsonFailedResponse(c, stateCode.NotAdmin, nil)
		c.Abort()
		return
	}
	c.Next()

}
