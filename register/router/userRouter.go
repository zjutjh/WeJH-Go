package router

import (
	"wejh-go/app/controllers/userController"

	"github.com/gin-gonic/gin"
	midsession "github.com/zjutjh/mygo/session/middleware"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/create/student/wechat", userController.BindOrCreateStudentUserFromWechat)
		user.POST("/create/student", userController.CreateStudentUser)

		user.POST("/login/wechat", userController.WeChatLogin)
		user.POST("/login", userController.AuthByPassword)
		user.POST("/login/session", userController.AuthBySession)

		user.POST("/info", midsession.Auth(), userController.GetUserInfo)

		user.POST("/del", midsession.Auth(), userController.DelAccount)
		user.POST("/repass", midsession.Auth(), userController.ResetPass)

		bind := user.Group("/bind", midsession.Auth())
		{
			bind.POST("/zf", userController.BindZFPassword)
			bind.POST("/yxy/send/code", userController.SendVerificationCode)
			bind.POST("/yxy/login", userController.LoginYxy)
			bind.POST("/oauth", userController.BindOauthPassword)
		}
	}
}
