package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

func GetUserInfo(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"studentID":  user.StudentID,
			"userType":   user.Type,
			"createTime": user.CreateTime,
		},
	})

}
