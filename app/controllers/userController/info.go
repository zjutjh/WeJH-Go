package userController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	//从 Redis/DB 获取最新的主题偏好
	theme, err := userServices.GetUserThemePreference(uint(user.ID))
	if err != nil {
		// 如果 Redis 和 DB 都查不到，进行降级处理：
		// 优先使用 session 里的旧值，如果没有则用默认的 "auto"
		theme = user.ThemePreference
		if theme == "" {
			theme = "auto"
		}
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
			"userType":        user.Type,
			"phoneNum":        user.PhoneNum,
			"createTime":      user.CreateTime,
			"themePreference": theme,
		},
	})

}
