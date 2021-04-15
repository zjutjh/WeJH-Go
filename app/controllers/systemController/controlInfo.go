package systemController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wejh-go/app/utils"
	"wejh-go/config/config"
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
