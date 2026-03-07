package router

import (
	"wejh-go/app/controllers/systemController"

	"github.com/gin-gonic/gin"
)

// 注册杂项路由
func systemRouterInit(r *gin.RouterGroup) {
	r.POST("/announcement", systemController.GetAnnouncement)
	r.POST("/applist", systemController.GetAppList)
	r.POST("/info", systemController.Info)
}
