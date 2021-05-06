package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/systemController"
)

// 注册杂项路由
func systemRouterInit(r *gin.RouterGroup) {
	r.Any("/announcement", systemController.GetAnnouncement)
	r.Any("/applist", systemController.GetAppList)
	r.Any("/info", systemController.Info)
}
