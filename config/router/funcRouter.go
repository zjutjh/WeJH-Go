package router

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/controllers/libraryController"
	"wejh-go/app/controllers/schoolCardController"
	"wejh-go/app/controllers/zfController"
	"wejh-go/app/midware"
)

func funcRouterInit(r *gin.RouterGroup) {
	user := r.Group("func")
	{
		bind := user.Group("/card", midware.CheckLogin)
		{
			bind.POST("/balance", schoolCardController.GetBalance)
			bind.POST("/history", schoolCardController.GetHistory)
			bind.POST("/today", schoolCardController.GetToday)
		}

		library := user.Group("/library", midware.CheckLogin)
		{
			library.POST("/current", libraryController.GetCurrent)
			library.POST("/history", libraryController.GetHistory)

		}

		zf := user.Group("/zf", midware.CheckLogin)
		{
			zf.POST("/classtable", zfController.GetClassTable)

		}

	}
}
