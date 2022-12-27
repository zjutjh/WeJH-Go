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

func SendVerificationCode(c *gin.Context) {
	var postForm phoneForm
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
	u := uuid.New()
	deviceId := u.String()
	userServices.SetDeviceID(user, deviceId)
	data, err := yxyServices.GetSecurityToken(deviceId)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if data.Level == 1 {
		img, err := yxyServices.GetCaptchaImage(deviceId, data.Token)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		utils.JsonResponse(c, apiException.YxyNeedCaptcha.Code, gin.H{
			"token": data,
			"img":   img,
		}, apiException.YxyNeedCaptcha.Msg)
		return
	}
	err = yxyServices.SendVerificationCode(data.Token, deviceId, "", postForm.PhoneNum)
	if err != nil && err != apiException.NotBindYxy {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
