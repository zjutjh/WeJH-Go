package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
)

type bindForm struct {
	PassWord string `json:"password"`
}

func BindZFPassword(c *gin.Context) {
	var postForm bindForm
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
	err = userServices.SetZFPassword(user, postForm.PassWord)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
func BindLibraryPassword(c *gin.Context) {
	var postForm bindForm
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
	err = userServices.SetLibraryPassword(user, postForm.PassWord)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func BindSchoolCardPassword(c *gin.Context) {
	var postForm bindForm
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
	err = userServices.SetCardPassword(user, postForm.PassWord)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
