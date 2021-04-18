package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/adminController"
	"wejh-go/app/midware"
)

// 注册杂项路由
func adminRouterInit(r *gin.RouterGroup) {

	admin := r.Group("admin", midware.CheckAdmin)
	{
		admin.POST("/announcement/create", adminController.CreateAnnouncement)
		admin.POST("/announcement/delete", adminController.DeleteAnnouncement)
		admin.POST("/announcement/update", adminController.UpdateAnnouncement)
	}
}
