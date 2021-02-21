package routers

import (
	"github.com/gin-gonic/gin"
	"wejh-go/controllers/misc_controllers"
)

// TODO: 在这个包下面建立一个 init 函数，日后方便初始化
// 注册杂项路由
func MiscRoutersInit(r *gin.Engine) {
	const pre = "/api" // TODO: 由于历史遗留原因，现在的所有接口要加上这个前缀，以后删掉
	r.GET(pre+"/time", misc_controllers.TimeController)
	r.GET(pre+"/announcement", misc_controllers.AnnouncementController)
	r.GET(pre+"/app-list", misc_controllers.AppListController)
}
