package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/adminController"
)

func initRouterInit(r *gin.RouterGroup) {
	set := r.Group("/set")
	{
		set.Any("/init", adminController.SetInit)
		set.POST("/encrypt", adminController.SetEncryptKey)
		set.POST("/terminfo", adminController.SetTermInfo)
	}
}
