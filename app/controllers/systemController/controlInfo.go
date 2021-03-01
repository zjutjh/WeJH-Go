package systemController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/config"
	"wejh-go/exception"
)

func GetControlInfo(c *gin.Context) {

}

func getTermInfo(c *gin.Context) {
	startTimeString := config.Config.GetString("term.term_start_date")
	currentTerm := config.Config.GetString("term.current_term") // 当前学期
	startTime, _ := time.Parse("2006-01-02", startTimeString)   // 学期开始的时间
	currentTime := time.Now()

	// 计算第几周
	week := (currentTime.Unix()-startTime.Unix())/3600/24/7 + 1

	utils.JsonSuccessResponse(c, gin.H{
		"day":      time.Now().Weekday(),
		"is_begin": week > 0,
		"month":    fmt.Sprintf("%d", time.Now().Month()),
		"term":     currentTerm,
		"week":     week,
	})

}

func getAppList(c *gin.Context) {
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
