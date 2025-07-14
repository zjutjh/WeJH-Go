package userController

import (
	"errors"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/config/wechat"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type autoLoginForm struct {
	Code string `json:"code" binding:"required"`
}
type passwordLoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthByPassword(c *gin.Context) {
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	user, err := userServices.GetUserByUsername(postForm.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.UserNotFind, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	if user.Type != models.Postgraduate && user.Type != models.Undergraduate {
		err = userServices.CheckLocalLogin(user, postForm.Password)
	} else {
		err = userServices.CheckLogin(postForm.Username, postForm.Password)
	}
	if err != nil {
		apiException.AbortWithError(c, err)
		return
	}

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf": user.ZFPassword != "",
				//"lib":   user.LibPassword != "",
				"yxy":   user.YxyUid != "",
				"oauth": user.OauthPassword != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}

func AuthBySession(c *gin.Context) {
	user, err := sessionServices.UpdateUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf": user.ZFPassword != "",
				//"lib": user.LibPassword != "" //已废弃
				"yxy": user.YxyUid != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}

func WeChatLogin(c *gin.Context) {
	var postForm autoLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)
	if err != nil {
		apiException.AbortWithException(c, apiException.OpenIDError, err)
		return
	}

	user := userServices.GetUserByWechatOpenID(session.OpenID)
	if user == nil {
		apiException.AbortWithException(c, apiException.UserNotFind, err)
		return
	}

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf": user.ZFPassword != "",
				//"lib":   user.LibPassword != "",
				"yxy":   user.YxyUid != "",
				"oauth": user.OauthPassword != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}
