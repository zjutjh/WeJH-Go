package schoolBusController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/schoolBusSearchRecordServices"
	"wejh-go/app/services/schoolBusServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

type FetchSchoolBusForm struct {
	Date        string `json:"date"`
	Departure   string `json:"departure"`
	Destination string `json:"destination"`
	StartTime   string `json:"startTime"`
}

var cstZone = time.FixedZone("GMT", 8*3600)

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
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	var busType models.SchoolBusType
	startDate, err := time.Parse("2006-01-02", postForm.Date)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	if startDate.Weekday() == time.Sunday || startDate.Weekday() == time.Saturday {
		busType = models.Weekend
	} else {
		busType = models.Weekday
	}
	_, err = time.ParseInLocation("15:04:05", postForm.StartTime, cstZone)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	schoolBuses, err := schoolBusServices.GetSchoolBus(models.SchoolBus{
		Departure:   postForm.Departure,
		Destination: postForm.Destination,
		StartTime:   postForm.StartTime,
		Type:        busType,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	err = SubmitRecord(models.SchoolBusSearchRecord{
		Username:    user.Username,
		Departure:   postForm.Departure,
		Destination: postForm.Destination,
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

func RecommendBus(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	record, err := schoolBusSearchRecordServices.GetRecord(user.Username)
	if err == gorm.ErrRecordNotFound {
		utils.JsonSuccessResponse(c, nil)
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	timeRaw := time.Now()
	timeNow := timeRaw.Format("15:04:05")
	var busType models.SchoolBusType
	if timeRaw.Weekday() == time.Sunday || timeRaw.Weekday() == time.Saturday {
		busType = models.Weekend
	} else {
		busType = models.Weekday
	}
	bus, err := schoolBusServices.RecommendSchoolBus(models.SchoolBus{
		Departure:   record.Departure,
		Destination: record.Destination,
		StartTime:   timeNow,
		Type:        busType,
	})
	fmt.Println(bus)
	if err == gorm.ErrRecordNotFound {
		bus, err = schoolBusServices.RecommendSchoolBus(models.SchoolBus{
			Departure:   record.Departure,
			Destination: record.Destination,
			StartTime:   "00:00",
			Type:        busType,
		})
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	busOpposite, err := schoolBusServices.RecommendSchoolBus(models.SchoolBus{
		Departure:   record.Destination,
		Destination: record.Departure,
		StartTime:   timeNow,
		Type:        busType,
	})
	if err == gorm.ErrRecordNotFound {
		busOpposite, err = schoolBusServices.RecommendSchoolBus(models.SchoolBus{
			Departure:   record.Destination,
			Destination: record.Departure,
			StartTime:   "00:00",
			Type:        busType,
		})
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	bus = append(bus, busOpposite...)
	fmt.Println(bus)

	utils.JsonSuccessResponse(c, bus)
}
