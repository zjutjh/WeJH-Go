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

	balance, err := funnelServices.GetClassTable(user, postForm.Year, postForm.Term)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

func GetScore(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		return
	}

	balance, err := funnelServices.GetScore(user, postForm.Year, postForm.Term)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

func GetExam(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		return
	}

	balance, err := funnelServices.GetExam(user, postForm.Year, postForm.Term)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}
