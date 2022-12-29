package schoolCardController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"
)

type recordForm struct {
}

func GetBalance(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	balance, err := yxyServices.GetCardBalance(user.DeviceID, user.YxyUid)
	if err == apiException.YxySessionExpired {
		_ = c.AbortWithError(200, apiException.YxySessionExpired)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

func GetRecord(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	balance, err := yxyServices.GetCardBalance(user.DeviceID, user.YxyUid)
	if err == apiException.YxySessionExpired {
		_ = c.AbortWithError(200, apiException.YxySessionExpired)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}
