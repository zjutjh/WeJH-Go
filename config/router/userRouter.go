package router

import (
	"wejh-go/app/controllers/userController"
	"wejh-go/app/midwares"

	"github.com/gin-gonic/gin"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/create/student/wechat", userController.BindOrCreateStudentUserFromWechat)
		user.POST("/create/student", userController.CreateStudentUser)

		user.POST("/login/wechat", userController.WeChatLogin)
		user.POST("/login", userController.AuthByPassword)
		user.POST("/login/session", userController.AuthBySession)

		user.POST("/info", midwares.CheckLogin, userController.GetUserInfo)

		user.POST("/del", midwares.CheckLogin, userController.DelAccount)
		user.POST("/repass", midwares.CheckLogin, userController.ResetPass)

		bind := user.Group("/bind", midwares.CheckLogin)
		{
			bind.POST("/zf", userController.BindZFPassword)
			bind.POST("/library", userController.BindLibraryPassword)
			bind.POST("/yxy/send/verification", userController.SendVerificationCode)
			bind.POST("/yxy/send/captcha", userController.SendVerificationCodeByCaptcha)
			bind.POST("/yxy/get/captcha", userController.GetCaptcha)
			bind.POST("/yxy/login", userController.LoginYxy)
			bind.POST("/oauth", userController.BindOauthPassword)
		}
	}
}
