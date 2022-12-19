package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/adminController"
	"wejh-go/app/midwares"
)

func initRouterInit(r *gin.RouterGroup) {
	set := r.Group("/set", midwares.CheckAdmin)
	{
		set.GET("/init", adminController.SetInit)
		set.POST("/encrypt", adminController.SetEncryptKey)
		set.POST("/terminfo", adminController.SetTermInfo)
	}
}
