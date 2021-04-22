package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/userController"
	"wejh-go/app/midware"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.POST("/create/student/wechat", userController.BindOrCreateStudentUserFromWechat)
		user.POST("/create/student", userController.CreateStudentUser)
		user.POST("/login/wechat",
			userController.WeChatLogin,
		)
		user.POST("/login",
			userController.AuthByPassword,
		)
		user.Any("/info",
			midware.CheckLogin, userController.GetUserInfo,
		)
		bind := user.Group("/bind", midware.CheckLogin)
		{
			bind.POST("/zf", userController.BindZFPassword)
			bind.POST("/library", userController.BindLibraryPassword)
			bind.POST("/schoolcard", userController.BindSchoolCardPassword)
		}

	}
}
