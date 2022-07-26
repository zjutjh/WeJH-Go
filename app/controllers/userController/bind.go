package userController

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L./lib -lyxy
#include <stdlib.h>
#include "./lib/yxy.h"
*/
import "C"

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

type phoneNumForm struct {
	PhoneNum string `json:"phoneNum"`
}

type verificationForm struct {
	VerificationCode string `json:"verificationCode"`
	DeviceId         string `json:"deviceId"`
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

func BindPhoneNum(c *gin.Context) {
	var phoneNumForm phoneNumForm
	err := c.ShouldBindJSON(&phoneNumForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	userServices.SetPhoneNum(user, phoneNumForm.PhoneNum)

	var loginHandle C.login_handle
	var securityTokenResult *C.security_token_result
	loginHandle.phone_num = C.CString(phoneNumForm.PhoneNum)
	C.gen_device_id(&loginHandle)
	errCode := C.get_security_token(&loginHandle, &securityTokenResult)
	if errCode != 0 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if securityTokenResult.level == 0 {
		errCode = C.send_verification_code(&loginHandle, securityTokenResult.token, nil)
		if errCode != 0 {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	} else {
		var imageResult C.CString
		errCode = C.get_captcha_image(&loginHandle, securityTokenResult.token, &imageResult)
		if errCode != 0 {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		errCode = C.send_verification_code(&loginHandle, securityTokenResult.token, imageResult)
		if errCode != 0 {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		C.free_c_string(imageResult)
	}

	C.free_security_token_result(securityTokenResult)
	utils.JsonSuccessResponse(c, gin.H{
		"deviceId": C.GoString(loginHandle.device_id),
	})
	C.free_c_string(loginHandle.device_id)
}

func BindYxyUid(c *gin.Context) {
	var verificationForm verificationForm
	err := c.ShouldBindJSON(&verificationForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	var loginHandle C.login_handle
	var loginResult *C.login_result
	loginHandle.phone_num = C.CString(user.PhoneNum)
	loginHandle.device_id = C.CString(verificationForm.DeviceId)
	errCode := C.do_login(&loginHandle, C.CString(verificationForm.VerificationCode), &loginResult)
	if errCode != 206 {
		_ = c.AbortWithError(200, apiException.WrongVerificationCode)
		return
	} else if errCode != 0 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	userServices.SetYxyUid(user, C.GoString(loginResult.uid))
	C.free_login_result(loginResult)
	utils.JsonSuccessResponse(c, nil)
}
