package schoolBusController

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/schoolBusSearchRecordServices"
	"wejh-go/app/services/schoolBusServices"
	"wejh-go/app/utils"
)

type FetchSchoolBusForm struct {
	Username  string `json:"username"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	From      string `json:"from"`
	To        string `json:"to"`
}

func GetBusList(c *gin.Context) {
	busList, err := schoolBusServices.GetSchoolBusList()
	if err != nil {
		_ = c.AbortWithError(200, err)
	} else {
		utils.JsonSuccessResponse(c, busList)
	}
}

func GetBus(c *gin.Context) {
	var postForm FetchSchoolBusForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	var tpyeNum models.SchoolBusType
	startDate, err := time.Parse("2006-01-02", postForm.Date)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	if startDate.Weekday() == time.Sunday || startDate.Weekday() == time.Saturday {
		tpyeNum = models.Weekend
	} else {
		tpyeNum = models.Weekday
	}
	startTime, err := time.Parse("15:04:05", postForm.StartTime)
	schoolBuses, err := schoolBusServices.GetSchoolBus(models.SchoolBus{
		From:      postForm.From,
		To:        postForm.To,
		StartTime: startTime,
		Type:      tpyeNum,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	err = SubmitRecord(models.SchoolBusSearchRecord{
		Username: postForm.Username,
		From:     postForm.From,
		To:       postForm.To,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, schoolBuses)
}

func SubmitRecord(record models.SchoolBusSearchRecord) error {
	result, err := schoolBusSearchRecordServices.GetRecord(record.Username)
	if err == gorm.ErrRecordNotFound {
		err := schoolBusSearchRecordServices.CreateRecord(record)
		return err
	} else if err != nil {
		return err
	}
	err = schoolBusSearchRecordServices.UpdateRecord(result.ID, record)
	return err
}
