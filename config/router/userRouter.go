package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/userController"
	"wejh-go/app/midware"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.POST("/autologin",
			userController.AutoLogin,
		)
		user.POST("/login",
			userController.AuthByPassword,
		)
		bind := user.Group("/bind", midware.CheckWechatSession)
		{
			bind.POST("/jh", userController.BindOrCreateUserFromWechat)
			bind.POST("/zf", userController.BindZFPassword)
			bind.POST("/library", userController.BindLibraryPassword)
			bind.POST("/schoolcard", userController.BindSchoolCardPassword)
		}

	}
}
