package systemController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/applistServices"
	"wejh-go/app/utils"
)

func GetAppList(c *gin.Context) {
	applists := applistServices.GetAppList(10)

	utils.JsonSuccessResponse(c, applists)
}
