package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/userController"
)

func userRouterInit(r *gin.RouterGroup) {

	r.POST("/code"+"/weapp", userController.WeAppController) // 返回用户的 openID
	r.POST("/login", userController.BindJHControllers)
	r.POST("/autoLogin", userController.AutoLoginControllers) // 自动登陆接口
	bindRoute := r.Group("/bind")
	bindRoute.POST("/jh", userController.BindJHControllers)
}
