package electricityController

import (
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

func GetBalance(c *gin.Context) {
	var postForm CampusForm
	err := c.ShouldBindQuery(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	token, err := yxyServices.GetElecAuthToken(user.YxyUid)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if postForm.Campus != "mgs" {
		postForm.Campus = "zhpf"
	}
	balance, err := yxyServices.ElectricityBalance(*token, postForm.Campus)
	if err == apiException.CampusMismatch {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, balance)
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
	token, err := yxyServices.GetElecAuthToken(user.YxyUid)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if postForm.Campus != "mgs" {
		postForm.Campus = "zhpf"
	}
	roomStrConcat, err := yxyServices.GetElecRoomStrConcat(*token, postForm.Campus, user.YxyUid)
	if err == apiException.CampusMismatch {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	records, err := yxyServices.ElectricityRechargeRecords(*token, postForm.Campus, postForm.Page, *roomStrConcat)
	if err == apiException.CampusMismatch {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, records.List)
}

func GetConsumptionRecords(c *gin.Context) {
	var postForm CampusForm
	err := c.ShouldBindQuery(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	token, err := yxyServices.GetElecAuthToken(user.YxyUid)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if postForm.Campus != "mgs" {
		postForm.Campus = "zhpf"
	}
	roomStrConcat, err := yxyServices.GetElecRoomStrConcat(*token, postForm.Campus, user.YxyUid)
	if err == apiException.CampusMismatch {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	records, err := yxyServices.ElectricityConsumptionRecords(*token, postForm.Campus, *roomStrConcat)
	if err == apiException.CampusMismatch {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, records.List)
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
