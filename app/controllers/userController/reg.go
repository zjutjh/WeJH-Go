package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/config/database"
	"wejh-go/config/wechat"
)

type createUserForm struct {
	UserName  string `json:"username"`
	PassWord  string `json:"password"`
	LoginType string `json:"type"`
}
type createUserWechatForm struct {
	UserName  string `json:"username"`
	PassWord  string `json:"password"`
	Code      string `json:"code"`
	LoginType string `json:"type"`
}

func BindOrCreateUserFromWechat(c *gin.Context) {
	var postForm createUserWechatForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)

	if err != nil {
		utils.JsonFailedResponse(c, stateCode.GetOpenIDFail, nil)
		return
	}

	err = userCenterServices.Auth(postForm.UserName, postForm.PassWord)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.UsernamePasswordUnmatched, nil)
		return
	}

	user, err := userServices.GetUserByOpenID(session.OpenID)
	if err != nil && user != nil {
		if user.JHPassword == postForm.PassWord {
			user.OpenID = session.OpenID
			database.DB.Save(user)
		}

	} else {
		user = &models.User{OpenID: session.OpenID, JHPassword: postForm.PassWord, Username: postForm.UserName}
		database.DB.Create(&user)
	}

	utils.JsonSuccessResponse(c, nil)
}

func CreateUser(c *gin.Context) {
	var postForm createUserForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}

	err = userCenterServices.Auth(postForm.UserName, postForm.PassWord)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.UsernamePasswordUnmatched, nil)
		return
	}

	user, err := userServices.GetUserByStudentID(postForm.UserName)
	if err != nil && user != nil {
		utils.JsonFailedResponse(c, stateCode.UserAlreadyExisted, nil)
		return
	}

	user = &models.User{JHPassword: postForm.PassWord, Username: postForm.UserName}
	database.DB.Create(&user)

	utils.JsonSuccessResponse(c, nil)
}
