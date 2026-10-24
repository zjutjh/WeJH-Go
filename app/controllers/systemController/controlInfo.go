package systemController

import (
	"time"
	"wejh-go/app/config"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	year, term, startTimeString, scoreYear, scoreTerm := config.GetTermInfo()
	startTime, _ := time.ParseInLocation("2006-01-02", startTimeString, time.Local) // 学期开始的时间
	currentTime := time.Now()
	schoolBusUrl := config.GetSchoolBusUrl()
	jpgUrl := config.GetWebpUrlKey()
	fileUrl := config.GetFileUrlKey()
	registerTips := config.GetRegisterTipsKey()

	daysToMonday := (int(time.Monday) - int(startTime.Weekday()) + 7) % 7
	weekStart := startTime.AddDate(0, 0, daysToMonday)
	today := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
	days := int(today.Sub(weekStart) / (24 * time.Hour))
	week := -1
	if days >= 0 {
		week = days/7 + 1
	}

	response := gin.H{
		"time":          currentTime,
		"is_begin":      week > 0,
		"termStartDate": startTimeString,
		"termYear":      year,
		"term":          term,
		"scoreYear":     scoreYear,
		"scoreTerm":     scoreTerm,
		"week":          week,
		"schoolBusUrl":  schoolBusUrl,
		"jpgUrl":        jpgUrl,
		"fileUrl":       fileUrl,
		"registerTips":  registerTips,
	}
	utils.JsonSuccessResponse(c, response)
}
