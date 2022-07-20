package adminController

import (
	"github.com/gin-gonic/gin"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/schoolBusServices"
	"wejh-go/app/utils"
)

type createSchoolBusForm struct {
	Line      int
	From      string
	To        string
	StartTime time.Time
}
type updateSchoolBusForm struct {
	ID        int `json:"id" binding:"required"`
	Line      int
	From      string
	To        string
	StartTime time.Time
}
type deleteSchoolBusForm struct {
	ID int `json:"id" binding:"required"`
}

func CreateSchoolBus(c *gin.Context) {
	var postForm createSchoolBusForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = schoolBusServices.CreateSchoolBus(models.SchoolBus{
		Line:      postForm.Line,
		From:      postForm.From,
		To:        postForm.To,
		StartTime: postForm.StartTime,
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

	err = schoolBusServices.UpdateSchoolBus(postForm.ID, models.SchoolBus{
		Line:      postForm.Line,
		From:      postForm.From,
		To:        postForm.To,
		StartTime: postForm.StartTime,
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
