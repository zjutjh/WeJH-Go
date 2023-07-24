package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/adminController"
	"wejh-go/app/controllers/funcControllers/lostAndFoundRecordController"
	"wejh-go/app/midwares"
)

// 注册杂项路由
func adminRouterInit(r *gin.RouterGroup) {

	admin := r.Group("/admin", midwares.CheckAdmin)
	{
		admin.GET("/check", adminController.CheckAdmin)
		announcement := admin.Group("/announcement")
		{
			announcement.POST("/create", adminController.CreateAnnouncement)
			announcement.POST("/delete", adminController.DeleteAnnouncement)
			announcement.POST("/update", adminController.UpdateAnnouncement)
		}
		applist := admin.Group("/applist")
		{
			applist.POST("/create", adminController.CreateApplist)
			applist.POST("/delete", adminController.DeleteApplist)
			applist.POST("/update", adminController.UpdateApplist)
		}
		schoolbus := admin.Group("/schoolbus")
		{
			schoolbus.POST("/create", adminController.CreateSchoolBus)
			schoolbus.POST("/delete", adminController.DeleteSchoolBus)
			schoolbus.POST("/update", adminController.UpdateSchoolBus)
		}
		set := admin.Group("/set")
		{
			set.GET("/reset", adminController.ResetInit)
			set.POST("/encrypt", adminController.SetEncryptKey)
			set.POST("/terminfo", adminController.SetTermInfo)
		}
		user := admin.Group("/user")
		{
			user.POST("/create", adminController.CreateAdminAccount)
		}
	}

	forU := r.Group("/foru", midwares.CheckLostAndFoundAdmin)
	{
		forU.POST("/lost", lostAndFoundRecordController.InsertRecord)
		forU.POST("/upload_img", lostAndFoundRecordController.UploadImg)
		forU.PUT("/lost", lostAndFoundRecordController.UpdateRecord)
		forU.GET("/lost", lostAndFoundRecordController.GetRecordsByAdmin)
		forU.DELETE("/lost", lostAndFoundRecordController.DeleteRecord)
	}
}
