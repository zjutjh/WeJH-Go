package systemController

import (
	"time"
	"wejh-go/app/config"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

var cst = time.FixedZone("CST", 8*3600)

func Info(c *gin.Context) {
	year, term, startTimeString, scoreYear, scoreTerm := config.GetTermInfo()
	startTime, _ := time.Parse("2006-01-02", startTimeString) // 学期开始的时间
	currentTime := time.Now()
	schoolBusUrl := config.GetSchoolBusUrl()
	jpgUrl := config.GetWebpUrlKey()
	fileUrl := config.GetFileUrlKey()
	registerTips := config.GetRegisterTipsKey()

	week := termWeek(currentTime, startTime)

	response := gin.H{
		"time":          time.Now(),
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

func weekBase(startDate time.Time) time.Time {
	t := startDate.In(cst)
	startDate1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if startDate1.Weekday() == time.Monday {
		return startDate1
	}
	return startDate1.AddDate(0, 0, (8-int(startDate1.Weekday()))%7)

}

func termWeek(now, startDate time.Time) int {
	t := now.In(cst)
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	days := int(today.Sub(weekBase(startDate)) / (24 * time.Hour))
	if days < 0 {
		return -1
	}
	return days/7 + 1

}
