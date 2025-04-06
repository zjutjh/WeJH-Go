package schoolCardController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
	if user.YxyUid == "" {
		_ = c.AbortWithError(200, apiException.NotBindYxy)
		return
	}
	token, err := yxyServices.GetCardAuthToken(user.YxyUid)
	if err == redis.Nil {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, false)
		_ = c.AbortWithError(200, apiException.YxySessionExpired)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	balance, err := yxyServices.GetCardBalance(user.DeviceID, user.YxyUid, user.PhoneNum, *token)
	if err == apiException.NotBindCard {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, true)
		_ = c.AbortWithError(200, err)
		return
	} else if err == apiException.YxySessionExpired {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, false)
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
	if user.YxyUid == "" {
		_ = c.AbortWithError(200, apiException.NotBindYxy)
		return
	}
	token, err := yxyServices.GetCardAuthToken(user.YxyUid)
	if err == redis.Nil {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, false)
		_ = c.AbortWithError(200, apiException.YxySessionExpired)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	records, err := yxyServices.GetCardConsumptionRecord(user.DeviceID, user.YxyUid, user.PhoneNum, *token, postForm.QueryTime)
	if err == apiException.NotBindCard {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, true)
		_ = c.AbortWithError(200, err)
		return
	} else if err == apiException.YxySessionExpired {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, false)
		_ = c.AbortWithError(200, err)
		return
	} else if err == apiException.ParamError {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, records.List)
}
