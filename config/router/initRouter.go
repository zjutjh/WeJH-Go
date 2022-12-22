package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/adminController"
)

func initRouterInit(r *gin.RouterGroup) {
	r.GET("/init", adminController.SetInit)
	r.POST("/encrypt", adminController.SetEncryptKey)
	r.POST("/terminfo", adminController.SetTermInfo)
}
