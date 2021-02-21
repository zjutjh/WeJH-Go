package routers

import (
	"github.com/gin-gonic/gin"
	"wejh-go/controllers"
)

func UserRoutersInit(r *gin.Engine) {
	const pre = "/api"                                        // TODO: 由于历史遗留原因加上了这样的前缀，以后删掉
	r.POST(pre+"/code"+"/weapp", controllers.WeAppController) // 返回用户的 openID
	r.POST(pre+"/login", controllers.BindJHControllers)
	r.POST(pre+"/autoLogin", controllers.AutoLoginControllers) // 自动登陆接口
	bindRoute := r.Group("/bind")
	bindRoute.POST("/jh", controllers.BindJHControllers)
}
