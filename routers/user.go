package routers

import (
	"github.com/gin-gonic/gin"
	"wejh-go/controllers"
)

func UserRoutersInit(r *gin.Engine) {
	const pre = "/api"                                  // 由于历史遗留原因加上了这样的前缀
	r.POST(pre+"/login", controllers.BindJHControllers) // 为了兼容
	bindRoute := r.Group("/bind")
	bindRoute.POST("/jh", controllers.BindJHControllers)
}
