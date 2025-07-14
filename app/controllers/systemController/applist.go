package systemController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/applistServices"
	"wejh-go/app/utils"
)

func GetAppList(c *gin.Context) {
	appLists, err := applistServices.GetAppList(10)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
	} else {
		utils.JsonSuccessResponse(c, appLists)
	}
}
