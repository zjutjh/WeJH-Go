package routers

import (
	"github.com/gin-gonic/gin"
	"wejh-go/controllers"
)

// 注册杂项路由
func MiscRoutersInit(r *gin.Engine) {
	const pre = "/api" // 由于历史遗留原因，现在的所有接口要加上这个前缀
	r.GET(pre+"/time", controllers.TimeController)
	r.GET(pre+"/announcement", controllers.AnnouncementController)
}
