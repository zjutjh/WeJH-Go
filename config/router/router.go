package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/midwares"
)

func Init(r *gin.Engine) {

	const pre = "/api"

	api := r.Group(pre, midwares.CheckInit)
	{
		systemRouterInit(api)
		userRouterInit(api)
		funcRouterInit(api)
		adminRouterInit(api)
	}
	init := r.Group(pre, midwares.CheckUninit)
	{
		initRouterInit(init)
	}
}
