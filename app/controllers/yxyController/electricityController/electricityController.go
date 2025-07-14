package electricityController

import (
	"errors"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

type recordForm struct {
	Page   string `form:"page" json:"page"`
	Campus string `form:"campus" json:"campus"`
}

type CampusForm struct {
	Campus string `form:"campus"`
}

type SubscribeLowBatteryAlertReq struct {
	Campus    string `json:"campus"`
	Threshold int    `json:"threshold"`
}

func GetBalance(c *gin.Context) {
	var postForm CampusForm
	err := c.ShouldBindQuery(&postForm)
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
	token, err := yxyServices.GetElecAuthToken(user.YxyUid)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	if postForm.Campus != "mgs" {
		postForm.Campus = "zhpf"
	}
	balance, err := yxyServices.ElectricityBalance(*token, postForm.Campus)
	if errors.Is(err, apiException.NotBindCard) {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, true)
		apiException.AbortWithError(c, err)
		return
	} else if errors.Is(err, apiException.CampusMismatch) {
		apiException.AbortWithError(c, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

func GetRechargeRecords(c *gin.Context) {
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
	token, err := yxyServices.GetElecAuthToken(user.YxyUid)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	if postForm.Campus != "mgs" {
		postForm.Campus = "zhpf"
	}
	roomStrConcat, err := yxyServices.GetElecRoomStrConcat(*token, postForm.Campus, user.YxyUid)
	if errors.Is(err, apiException.NotBindCard) {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, true)
		apiException.AbortWithError(c, err)
		return
	} else if errors.Is(err, apiException.CampusMismatch) {
		apiException.AbortWithError(c, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	records, err := yxyServices.ElectricityRechargeRecords(*token, postForm.Campus, postForm.Page, *roomStrConcat)
	if errors.Is(err, apiException.CampusMismatch) {
		apiException.AbortWithError(c, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, records.List)
}

func GetConsumptionRecords(c *gin.Context) {
	var postForm CampusForm
	err := c.ShouldBindQuery(&postForm)
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
	token, err := yxyServices.GetElecAuthToken(user.YxyUid)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	if postForm.Campus != "mgs" {
		postForm.Campus = "zhpf"
	}
	roomStrConcat, err := yxyServices.GetElecRoomStrConcat(*token, postForm.Campus, user.YxyUid)
	if errors.Is(err, apiException.NotBindCard) {
		_ = yxyServices.Unbind(user.ID, user.YxyUid, true)
		apiException.AbortWithError(c, err)
		return
	} else if errors.Is(err, apiException.CampusMismatch) {
		apiException.AbortWithError(c, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	records, err := yxyServices.GetElecConsumptionRecords(*token, postForm.Campus, *roomStrConcat)
	if errors.Is(err, apiException.CampusMismatch) {
		apiException.AbortWithError(c, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, records.List)
}

func SubscribeLowBatteryAlert(c *gin.Context) {
	// var req SubscribeLowBatteryAlertReq
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	apiException.AbortWithException(c, apiException.ParamError)
	// 	return
	// }
	// 临时兼容用
	req := SubscribeLowBatteryAlertReq{
		Campus:    "zhpf",
		Threshold: 20,
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
	if req.Campus != "mgs" {
		req.Campus = "zhpf"
	}
	if req.Threshold <= 0 {
		req.Threshold = 20
	}
	if err := yxyServices.SubscribeLowBatteryAlert(models.LowBatteryAlertSubscription{
		UserID:    user.ID,
		Campus:    req.Campus,
		Threshold: req.Threshold,
	}); err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func GetLowBatteryAlertSubscription(c *gin.Context) {
	var form CampusForm
	err := c.ShouldBindQuery(&form)
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
	if form.Campus != "mgs" {
		form.Campus = "zhpf"
	}
	subscription, err := yxyServices.GetOrCreateLowBatteryAlertSubscription(user.ID, form.Campus)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, subscription)
}
