package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/userController"
	"wejh-go/app/midwares"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/create/student/wechat", userController.BindOrCreateStudentUserFromWechat)
		user.POST("/create/student", userController.CreateStudentUser)

		user.POST("/login/wechat", userController.WeChatLogin)
		user.POST("/login", userController.AuthByPassword)

		user.POST("/info", midwares.CheckLogin, userController.GetUserInfo)
		bind := user.Group("/bind", midwares.CheckLogin)
		{
			bind.POST("/zf", userController.BindZFPassword)
			bind.POST("/library", userController.BindLibraryPassword)
			bind.POST("/yxy/send/verification", userController.SendVerificationCode)
			bind.POST("/yxy/send/captcha", userController.SendVerificationCodeByCaptcha)
			bind.POST("/yxy/get/captcha", userController.GetCaptcha)
			bind.POST("/yxy/login", userController.LoginYxy)
		}
	}
}
