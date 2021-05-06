package zfController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
)

type form struct {
	Year string `json:"year" binding:"required"`
	Term string `json:"term" binding:"required"`
}

func GetClassTable(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}

	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		return
	}

	result, err := funnelServices.GetClassTable(user, postForm.Year, postForm.Term)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetScore(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		return
	}

	result, err := funnelServices.GetScore(user, postForm.Year, postForm.Term)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetExam(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		return
	}

	result, err := funnelServices.GetExam(user, postForm.Year, postForm.Term)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

type roomForm struct {
	Year     string `json:"year" binding:"required"`
	Term     string `json:"term" binding:"required"`
	Campus   string `json:"campus" binding:"required"`
	Weekday  string `json:"weekday" binding:"required"`
	Sections string `json:"sections" binding:"required"`
	Week     string `json:"week" binding:"required"`
}

func GetRoom(c *gin.Context) {
	var postForm roomForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		return
	}

	result, err := funnelServices.GetRoom(user, postForm.Year, postForm.Term, postForm.Campus, postForm.Weekday, postForm.Week, postForm.Sections)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}
