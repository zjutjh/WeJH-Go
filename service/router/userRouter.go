package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/userController"
	"wejh-go/app/midware"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.POST("/login",
			userController.AutoLogin,
		)
		bind := user.Group("/bind", midware.CheckWechatSession)
		{
			bind.POST("/jh", userController.BindJHID)
			bind.POST("/zf", userController.BindZFPassword)
			bind.POST("/library", userController.BindLibraryPassword)
			bind.POST("/schoolcard", userController.BindSchoolCardPassword)
		}

	}
}
