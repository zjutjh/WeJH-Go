package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/systemController"
)

// 注册杂项路由
func systemRouterInit(r *gin.RouterGroup) {
	r.GET("/announcement", systemController.GetAnnouncement)
	r.GET("/applist", systemController.GetAppList)
}
