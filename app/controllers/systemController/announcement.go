package systemController

import (
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/config"
	"wejh-go/exception"
)

func GetAnnouncement(c *gin.Context) {
	var announcement map[string]interface{}
	err := config.Config.UnmarshalKey("announcement", &announcement)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.SystemError, nil)
		log.Fatal(exception.ConfigError, err)
	}
	utils.JsonSuccessResponse(c, announcement)
}
