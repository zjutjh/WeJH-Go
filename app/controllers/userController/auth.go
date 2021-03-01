package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/service/wechat"
)

type autoLoginForm struct {
	Code      string `json:"code" binding:"required"`
	LoginType string `json:"type"`
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
			"studentID": user.StudentID, // TODO: 添加其他账号的绑定信息
		},
	})
}
