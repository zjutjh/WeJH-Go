package systemController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/applistServices"
	"wejh-go/app/utils"
)

func GetAppList(c *gin.Context) {
	applists, err := applistServices.GetAppList(10)
	if err != nil {
		utils.JsonErrorResponse(c, err)
	} else {
		utils.JsonSuccessResponse(c, applists)
	}
}
