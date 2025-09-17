package schoolCardController

import (
	"errors"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

type recordForm struct {
	QueryTime string `json:"queryTime"`
}

func GetBalance(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	if user.YxyUid == "" {
		apiException.AbortWithException(c, apiException.NotBindYxy, nil)
		return
	}
	balance, err := yxyServices.GetCardBalance(user.DeviceID, user.YxyUid, user.PhoneNum)
	if errors.Is(err, apiException.NotBindCard) {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, true)
		apiException.AbortWithError(c, err)
		return
	} else if errors.Is(err, apiException.YxySessionExpired) {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, false)
		apiException.AbortWithError(c, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

func GetConsumptionRecord(c *gin.Context) {
	var postForm recordForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	if user.YxyUid == "" {
		apiException.AbortWithException(c, apiException.NotBindYxy, nil)
		return
	}

	records, err := yxyServices.GetCardConsumptionRecord(user.DeviceID, user.YxyUid, user.PhoneNum, postForm.QueryTime)
	if errors.Is(err, apiException.NotBindCard) {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, true)
		apiException.AbortWithError(c, err)
		return
	} else if errors.Is(err, apiException.YxySessionExpired) {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, false)
		apiException.AbortWithError(c, err)
		return
	} else if errors.Is(err, apiException.ParamError) {
		apiException.AbortWithError(c, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, records.List)
}
