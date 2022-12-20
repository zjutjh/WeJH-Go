package userController

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
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
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := userServices.GetUserByUsername(postForm.Username)
	if err == gorm.ErrRecordNotFound {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		return
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	h := sha256.New()
	h.Write([]byte(postForm.Password))
	pass := hex.EncodeToString(h.Sum(nil))
	if user.JHPassword != pass {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf":   user.ZFPassword != "",
				"lib":  user.LibPassword != "",
				"card": user.CardPassword != "",
				"yxy":  user.YXYUid != "",
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
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)
	if err != nil {
		_ = c.AbortWithError(200, apiException.OpenIDError)
		return
	}

	user := userServices.GetUserByWechatOpenID(session.OpenID)
	if user == nil {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		return
	}

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf":   user.ZFPassword != "",
				"lib":  user.LibPassword != "",
				"card": false,
				"yxy":  user.YXYUid != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}
