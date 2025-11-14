package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/systemController"
)

// 注册杂项路由
func systemRouterInit(r *gin.RouterGroup) {
	r.POST("/announcement", systemController.GetAnnouncement)
	r.POST("/applist", systemController.GetAppList)
	r.POST("/info", systemController.Info)
}
