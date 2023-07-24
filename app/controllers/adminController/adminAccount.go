package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
)

type AccountForm struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	AdminType int    `json:"admintype"`
}

func CreateAdminAccount(c *gin.Context) {
	var postForm AccountForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	err = userServices.CreateAdmin(postForm.UserName, postForm.Password, postForm.AdminType)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
