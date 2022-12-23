package adminController

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/schoolBusServices"
	"wejh-go/app/utils"
)

type createSchoolBusForm struct {
	Line        string `json:"line" binding:"required"`
	Departure   string `json:"departure" binging:"required"`
	Destination string `json:"destination" binging:"required"`
	Type        int    `json:"type" binging:"required"`
	StartTime   string `json:"startTime" binging:"required"`
}
type updateSchoolBusForm struct {
	ID          int    `json:"id" binding:"required"`
	Line        string `json:"line" binding:"required"`
	Departure   string `json:"departure" binging:"required"`
	Destination string `json:"destination" binging:"required"`
	Type        int    `json:"type" binging:"required"`
	StartTime   string `json:"startTime" binging:"required"`
}
type deleteSchoolBusForm struct {
	ID int `json:"id" binding:"required"`
}

var cstZone = time.FixedZone("GMT", 8*3600)

func CreateSchoolBus(c *gin.Context) {
	var postForm createSchoolBusForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		log.Println(err.Error())
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	_, err = time.ParseInLocation("15:04:05", postForm.StartTime, cstZone)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	err = schoolBusServices.CreateSchoolBus(models.SchoolBus{
		Line:        postForm.Line,
		Departure:   postForm.Departure,
		Destination: postForm.Destination,
		Type:        models.SchoolBusType(postForm.Type),
		StartTime:   postForm.StartTime,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
func UpdateSchoolBus(c *gin.Context) {
	var postForm updateSchoolBusForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	_, err = time.ParseInLocation("15:04:05", postForm.StartTime, cstZone)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	println(postForm.StartTime)
	err = schoolBusServices.UpdateSchoolBus(postForm.ID, models.SchoolBus{
		Line:        postForm.Line,
		Departure:   postForm.Departure,
		Destination: postForm.Destination,
		Type:        models.SchoolBusType(postForm.Type),
		StartTime:   postForm.StartTime,
	},
	)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
func DeleteSchoolBus(c *gin.Context) {
	var postForm deleteSchoolBusForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = schoolBusServices.DeleteSchoolBus(postForm.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
