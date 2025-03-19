package userController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/circuitBreaker"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type bindForm struct {
	PassWord string `json:"password"`
}

type phoneForm struct {
	PhoneNum string `json:"phoneNum"`
}

type loginForm struct {
	PhoneNum string `json:"phoneNum"`
	Code     string `json:"code"`
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
	api, _, err := circuitBreaker.CB.GetApi(true, false)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	err = userServices.SetZFPassword(user, postForm.PassWord, api)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func BindOauthPassword(c *gin.Context) {
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
	api, _, err := circuitBreaker.CB.GetApi(false, true)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	err = userServices.SetOauthPassword(user, postForm.PassWord, api)
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
	_, err = sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	deviceId := uuid.New().String()
	data, err := yxyServices.GetSecurityToken(deviceId)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	err = yxyServices.SendVerificationCode(data.SecurityToken, deviceId, postForm.PhoneNum)
	if err == apiException.WrongPhoneNum || err == apiException.SendVerificationCodeLimit || err == apiException.NotBindYxy {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func LoginYxy(c *gin.Context) {
	var postForm loginForm
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
	deviceId := uuid.New().String()
	info, err := yxyServices.LoginByCode(postForm.Code, deviceId, postForm.PhoneNum)
	if err == apiException.WrongVerificationCode || err == apiException.WrongPhoneNum || err == apiException.NotBindYxy {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if err := yxyServices.SetCardAuthToken(info.UID, info.Token); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	userServices.SetDeviceID(user, deviceId)
	userServices.DecryptUserKeyInfo(user)
	userServices.SetYxyUid(user, info.UID)
	userServices.DecryptUserKeyInfo(user)
	userServices.SetPhoneNum(user, postForm.PhoneNum)
	utils.JsonSuccessResponse(c, nil)
}
