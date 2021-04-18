package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/funcControllers/canteenController"
	"wejh-go/app/controllers/funcControllers/libraryController"
	"wejh-go/app/controllers/funcControllers/schoolBusController"
	"wejh-go/app/controllers/funcControllers/schoolCardController"
	"wejh-go/app/controllers/funcControllers/zfController"
	"wejh-go/app/midware"
)

func funcRouterInit(r *gin.RouterGroup) {
	fun := r.Group("func")
	{
		fun.POST("/canteen/flow", canteenController.GetCanteenFlowRate)
		fun.POST("/bus", schoolBusController.GetBusList)

		card := fun.Group("/card", midware.CheckLogin)
		{
			card.POST("/balance", schoolCardController.GetBalance)
			card.POST("/history", schoolCardController.GetHistory)
			card.POST("/today", schoolCardController.GetToday)
		}

		library := fun.Group("/library", midware.CheckLogin)
		{
			library.POST("/current", libraryController.GetCurrent)
			library.POST("/history", libraryController.GetHistory)

		}

		zf := fun.Group("/zf", midware.CheckLogin)
		{
			zf.POST("/classtable", zfController.GetClassTable)
			zf.POST("/exam", zfController.GetExam)
			zf.POST("/score", zfController.GetScore)

		}

	}
}
