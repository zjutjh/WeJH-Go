package electricityController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"
)

type recordForm struct {
	Page string `json:"page"`
}

func GetBalance(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	token, err := yxyServices.Auth(user.YxyUid)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	electricityBalance, err := yxyServices.ElectricityBalance(*token)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, electricityBalance)
}

func GetRechargeRecords(c *gin.Context) {
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
	token, err := yxyServices.Auth(user.YxyUid)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	roomInfo, err := yxyServices.Bind(*token)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	records, err := yxyServices.ElectricityRechargeRecords(
		*token,
		(*roomInfo)["area_id"].(string),
		(*roomInfo)["building_code"].(string),
		(*roomInfo)["floor_code"].(string),
		(*roomInfo)["room_code"].(string),
		postForm.Page)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, records)
}

func GetConsumptionRecords(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	token, err := yxyServices.Auth(user.YxyUid)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	roomInfo, err := yxyServices.Bind(*token)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	electricityBalance, err := yxyServices.ElectricityBalance(*token)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	records, err := yxyServices.ElectricityConsumptionRecords(
		*token,
		(*roomInfo)["area_id"].(string),
		(*roomInfo)["building_code"].(string),
		(*roomInfo)["floor_code"].(string),
		(*roomInfo)["room_code"].(string),
		(*electricityBalance)["md_type"].(string))
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, records)
}

func InsertLowBatteryQuery(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.YxyUid == "" {
		_ = c.AbortWithError(200, apiException.NotBindYxy)
		return
	}
	err = yxyServices.InsertRecord(models.LowBatteryQueryRecord{
		UserID: user.ID,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
