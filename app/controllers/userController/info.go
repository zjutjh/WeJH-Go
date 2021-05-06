package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
)

func GetUserInfo(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.NotLogin, nil)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"userType":  user.Type,
			"bind": gin.H{
				"zf":   user.ZFPassword != "",
				"lib":  user.LibPassword != "",
				"card": user.CardPassword != "",
			},
			"createTime": user.CreateTime,
		},
	})

}
