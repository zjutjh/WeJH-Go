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
		_ = c.AbortWithError(200, apiException.ServerError)
	} else {
		utils.JsonSuccessResponse(c, appLists)
	}
}
