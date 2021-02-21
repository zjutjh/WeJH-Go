package systemController

import (
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/config"
	"wejh-go/exception"
)

func GetAppList(c *gin.Context) {
	var appList []map[string]interface{}
	var icons map[string]interface{}
	err := config.Config.UnmarshalKey("applicationsList", &appList)
	if err == nil {
		err = config.Config.UnmarshalKey("icons", &icons)
	}

	if err != nil {
		utils.JsonFailedResponse(c, stateCode.SystemError, nil)
		log.Fatal(exception.ConfigError, err)
	}

	utils.JsonSuccessResponse(c, gin.H{
		"app-list": appList,
		"icons":    icons,
	})

}
