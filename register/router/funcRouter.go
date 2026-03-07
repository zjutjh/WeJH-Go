package router

import (
	"wejh-go/app/controllers/funcControllers/customizeHomeController"
	"wejh-go/app/controllers/funcControllers/lessonController"
	"wejh-go/app/controllers/funcControllers/libraryController"
	"wejh-go/app/controllers/funcControllers/lostAndFoundRecordController"
	"wejh-go/app/controllers/funcControllers/noticeController"
	"wejh-go/app/controllers/funcControllers/suppliesController"
	"wejh-go/app/controllers/funcControllers/themeController"
	"wejh-go/app/controllers/funcControllers/zfController"
	"wejh-go/app/controllers/yxyController/busController"
	"wejh-go/app/controllers/yxyController/electricityController"
	"wejh-go/app/controllers/yxyController/schoolCardController"

	"github.com/gin-gonic/gin"
	midsession "github.com/zjutjh/mygo/session/middleware"
)

func funcRouterInit(r *gin.RouterGroup) {
	fun := r.Group("/func")
	{
		customizeHome := fun.Group("/home", midsession.Auth())
		{
			customizeHome.GET("/get", customizeHomeController.GetCustomizeHome)
			customizeHome.POST("/update", customizeHomeController.UpdateCustomizeHome)
		}

		lesson := fun.Group("/lesson", midsession.Auth())
		{
			lesson.POST("/create", lessonController.CreateLesson)
			lesson.POST("/get", lessonController.GetLesson)
			lesson.POST("/update", lessonController.UpdateLesson)
			lesson.POST("/delete", lessonController.DeleteLesson)
		}

		electricity := fun.Group("/electricity", midsession.Auth())
		{
			electricity.GET("/balance", electricityController.GetBalance)
			electricity.POST("/record", electricityController.GetRechargeRecords)
			electricity.GET("/consumption", electricityController.GetConsumptionRecords)
			electricity.GET("/subscription", electricityController.GetLowBatteryAlertSubscription)
			electricity.POST("/subscription", electricityController.SubscribeLowBatteryAlert)
		}

		bus := fun.Group("/bus", midsession.Auth())
		{
			bus.GET("/info", busController.GetBusInfo)
			bus.GET("/announcement", busController.GetBusAnnouncement)
			bus.GET("/config", busController.GetBusConfig)
		}

		card := fun.Group("/card", midsession.Auth())
		{
			card.GET("/balance", schoolCardController.GetBalance)
			card.POST("/record", schoolCardController.GetConsumptionRecord)
		}

		library := fun.Group("/library", midsession.Auth())
		{
			library.POST("/current", libraryController.GetCurrent)
			library.POST("/history", libraryController.GetHistory)
		}

		zf := fun.Group("/zf", midsession.Auth())
		{
			zf.POST("/classtable", zfController.GetClassTable)
			zf.POST("/exam", zfController.GetExam)
			zf.POST("/room", zfController.GetRoom)
			zf.POST("/score", zfController.GetScore)
			zf.POST("/midtermscore", zfController.GetMidTermScore)
		}

		lost := fun.Group("/lost", midsession.Auth())
		{
			lost.GET("", lostAndFoundRecordController.GetRecords)
			lost.GET("/kind_list", lostAndFoundRecordController.GetKindList)
		}

		notice := fun.Group("/information", midsession.Auth())
		{
			notice.GET("", noticeController.GetNotice)
		}

		// 正装借用
		suppliesBorrow := fun.Group("/supplies-borrow", midsession.Auth())
		{
			suppliesBorrow.GET("/qa", suppliesController.GetQAList)
			suppliesBorrow.GET("/supplies", suppliesController.GetSuppliesList)
			info := suppliesBorrow.Group("/info")
			{
				info.GET("", suppliesController.GetPersonalInfo)
				info.POST("", suppliesController.SavePersonalInfo)
			}
			borrow := suppliesBorrow.Group("/borrow")
			{
				borrow.POST("", suppliesController.ApplyBorrow)
				borrow.GET("", suppliesController.GetBorrowRecord)
				borrow.DELETE("", suppliesController.CancelBorrow)
			}
		}

		// 主题色
		theme := fun.Group("/theme", midsession.Auth())
		{
			theme.GET("/get", themeController.GetThemeList)
			theme.POST("/choose", themeController.ChooseCurrentTheme)
		}
	}
}
