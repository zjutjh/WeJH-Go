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

type captchaForm struct {
	Captcha  string `json:"captcha"`
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
	if yxyServices.CheckToken("SecurityToken" + user.Username) {
		yxyServices.DelToken("SecurityToken" + user.Username)
	}
	yxyServices.SetToken("SecurityToken"+user.Username, data.Token)
	if data.Level == 1 {
		_ = c.AbortWithError(200, apiException.YxyNeedCaptcha)
		return
	}
	err = yxyServices.SendVerificationCode(data.Token, deviceId, "", postForm.PhoneNum)
	if err == apiException.WrongCaptcha || err == apiException.NotBindYxy {
		_ = c.AbortWithError(200, err)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func GetCaptcha(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	token, err := yxyServices.GetToken("SecurityToken" + user.Username)
	if err != nil {
		data, err := yxyServices.GetSecurityToken(user.DeviceID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		token = &data.Token
		yxyServices.DelToken("SecurityToken" + user.Username)
		yxyServices.SetToken("SecurityToken"+user.Username, data.Token)
	}
	img, err := yxyServices.GetCaptchaImage(user.DeviceID, *token)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, img)
}

func SendVerificationCodeByCaptcha(c *gin.Context) {
	var postForm captchaForm
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
	token, err := yxyServices.GetToken("SecurityToken" + user.Username)
	if err != nil {
		_ = c.AbortWithError(200, apiException.YxyNeedCaptcha)
		return
	}
	err = yxyServices.SendVerificationCode(*token, user.DeviceID, postForm.Captcha, postForm.PhoneNum)
	if err == apiException.WrongCaptcha || err == apiException.NotBindYxy {
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
	uid, err := yxyServices.LoginByCode(postForm.Code, user.DeviceID, postForm.PhoneNum)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	userServices.SetYxyUid(user, *uid)
	userServices.SetPhoneNum(user, postForm.PhoneNum)
	utils.JsonSuccessResponse(c, nil)
}
