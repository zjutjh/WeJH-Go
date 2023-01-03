package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/funcControllers/canteenController"
	"wejh-go/app/controllers/funcControllers/lessonController"
	"wejh-go/app/controllers/funcControllers/libraryController"
	"wejh-go/app/controllers/funcControllers/schoolBusController"
	"wejh-go/app/controllers/funcControllers/zfController"
	"wejh-go/app/controllers/yxyController/electricityController"
	"wejh-go/app/controllers/yxyController/schoolCardController"
	"wejh-go/app/midwares"
)

func funcRouterInit(r *gin.RouterGroup) {
	fun := r.Group("/func")
	{
		// TODO 准备删除
		fun.POST("/canteen/flow", canteenController.GetCanteenFlowRate)

		lesson := fun.Group("/lesson", midwares.CheckLogin)
		{
			lesson.POST("/create", lessonController.CreateLesson)
			lesson.POST("/get", lessonController.GetLesson)
			lesson.POST("/update", lessonController.UpdateLesson)
			lesson.POST("/delete", lessonController.DeleteLesson)
		}

		electricity := fun.Group("/electricity", midwares.CheckLogin)
		{
			electricity.GET("/balance", electricityController.GetBalance)
			electricity.POST("/record", electricityController.GetRechargeRecords)
			electricity.GET("/consumption", electricityController.GetConsumptionRecords)
		}

		bus := fun.Group("/bus", midwares.CheckLogin)
		{
			bus.GET("/list", schoolBusController.GetBusList)
			bus.POST("/get", schoolBusController.GetBus)
			bus.GET("/recommend", schoolBusController.RecommendBus)
			bus.POST("/time", schoolBusController.GetTimeList)
		}

		card := fun.Group("/card", midwares.CheckLogin)
		{
			card.GET("/balance", schoolCardController.GetBalance)
			card.POST("/record", schoolCardController.GetConsumptionRecord)
		}

		library := fun.Group("/library", midwares.CheckLogin)
		{
			library.POST("/current", libraryController.GetCurrent)
			library.POST("/history", libraryController.GetHistory)

		}

		zf := fun.Group("/zf", midwares.CheckLogin)
		{
			zf.POST("/classtable", zfController.GetClassTable)
			zf.POST("/exam", zfController.GetExam)
			zf.POST("/room", zfController.GetRoom)
			zf.POST("/score", zfController.GetScore)
			zf.POST("/midtermscore", zfController.GetMidTermScore)
		}
	}
}
