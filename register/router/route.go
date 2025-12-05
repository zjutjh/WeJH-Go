package router

import (
	"wejh-go/app/midwares"

	"github.com/gin-gonic/gin"
	"github.com/zjutjh/mygo/middleware/cors"
	"github.com/zjutjh/mygo/session"
)

func Route(router *gin.Engine) {
	router.Use(cors.Pick())
	router.Use(session.Pick())
	router.Use(midwares.ErrHandler())
	router.NoMethod(midwares.HandleNotFound)
	router.NoRoute(midwares.HandleNotFound)
	api := router.Group(routePrefix(), midwares.CheckInit)
	{
		systemRouterInit(api)
		userRouterInit(api)
		funcRouterInit(api)
		adminRouterInit(api)
	}
	init := router.Group(routePrefix(), midwares.CheckUninit)
	{
		initRouterInit(init)
	}
}

func routePrefix() string {
	return "/api"
}
