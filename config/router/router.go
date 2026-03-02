package router

import (
	"wejh-go/app/midwares"

	"github.com/gin-gonic/gin"
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
