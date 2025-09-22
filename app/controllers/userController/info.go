package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

func GetUserInfo(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
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
