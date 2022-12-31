package schoolCardController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"
)

type recordForm struct {
	QueryTime string `json:"queryTime"`
}

func GetBalance(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	balance, err := yxyServices.GetCardBalance(user.DeviceID, user.YxyUid)
	if err == apiException.YxySessionExpired {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

func GetConsumptionRecord(c *gin.Context) {
	var postForm recordForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	balance, err := yxyServices.GetCardConsumptionRecord(user.DeviceID, user.YxyUid, postForm.QueryTime)
	if err == apiException.YxySessionExpired ||
		err == apiException.ParamError {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}
