package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"
)

type bindForm struct {
	PassWord string `json:"password"`
}

type phoneForm struct {
	PhoneNum string `json:"phoneNum"`
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

func BindYxy(c *gin.Context) {
	var postForm phoneForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	//user, err := sessionServices.GetUserSession(c)
	//if err != nil {
	//	_ = c.AbortWithError(200, apiException.NotLogin)
	//	return
	//}
	u := uuid.New()
	key := u.String()
	data, err := yxyServices.GetSecurityToken(key)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if data.Level == 1 {
		err = yxyServices.GetCaptchaImage(key, data.Token)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}

	utils.JsonSuccessResponse(c, data)
}
