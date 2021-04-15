package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/config/wechat"
)

type autoLoginForm struct {
	Code      string `json:"code" binding:"required"`
	LoginType string `json:"type"`
}
type passwordLoginForm struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	LoginType string `json:"type"`
}

func AuthByPassword(c *gin.Context) {
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm)

	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}
	user, err := userServices.GetUserByUsernameAndPassword(postForm.Username, postForm.Password)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}
	if user != nil {
		err = sessionServices.SetUserSession(c, user)
		utils.JsonSuccessResponse(c, gin.H{
			"user": gin.H{
				"ID": user.ID,
			},
		})
	}
}

func AutoLogin(c *gin.Context) {
	var postForm autoLoginForm
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

	user, err := userServices.GetUserByOpenID(session.OpenID)

	if err != nil {
		utils.JsonFailedResponse(c, stateCode.UserNotFind, nil)
		return
	}

	err = sessionServices.SetWechatSession(c, &session)

	if err != nil {
		utils.JsonFailedResponse(c, stateCode.SystemError, nil)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"ID": user.ID,
		},
	})
}
