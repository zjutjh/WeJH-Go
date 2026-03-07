package router

import (
	"wejh-go/app/controllers/adminController"

	"github.com/gin-gonic/gin"
)

func initRouterInit(r *gin.RouterGroup) {
	r.GET("/init", adminController.SetInit)
	r.POST("/encrypt", adminController.SetEncryptKey)
	r.POST("/systeminfo", adminController.SetSystemInfo)
}
