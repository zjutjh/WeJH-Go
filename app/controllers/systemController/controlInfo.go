package systemController

import (
	"github.com/gin-gonic/gin"
	"time"
	"wejh-go/app/config"
	"wejh-go/app/utils"
)

func Info(c *gin.Context) {
	year, term, startTimeString := config.GetTermInfo()
	startTime, _ := time.Parse("2006-01-02", startTimeString) // 学期开始的时间
	currentTime := time.Now()
	schoolBusUrl := config.GetSchoolBusUrl()

	week := ((currentTime.Unix()-startTime.Unix())/3600+8)/24/7 + 1
	if currentTime.Unix() < startTime.Unix()-8*3600 {
		week = -1
	}
	utils.JsonSuccessResponse(c, gin.H{
		"time":          time.Now(),
		"is_begin":      week > 0,
		"termStartDate": startTimeString,
		"termYear":      year,
		"term":          term,
		"week":          week,
		"score":         "下", //TODO 特殊时期过去后删除
		"schoolBusUrl":  schoolBusUrl,
	})

}
