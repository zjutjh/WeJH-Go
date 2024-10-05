package systemController

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/services/themeServices"
	"wejh-go/app/utils"
)

func Info(c *gin.Context) {
	year, term, startTimeString, scoreYear, scoreTerm := config.GetTermInfo()
	startTime, _ := time.Parse("2006-01-02", startTimeString) // 学期开始的时间
	currentTime := time.Now()
	schoolBusUrl := config.GetSchoolBusUrl()
	jpgUrl := config.GetWebpUrlKey()
	fileUrl := config.GetFileUrlKey()
	registerTips := config.GetRegisterTipsKey()
	defaultThemeIDStr := config.GetDefaultThemeKey()

	week := ((currentTime.Unix()-startTime.Unix())/3600+8)/24/7 + 1
	if currentTime.Unix() < startTime.Unix()-8*3600 {
		week = -1
	}

	var defaultTheme models.Theme
	if defaultThemeIDStr != "" {
		defaultThemeID, err := strconv.Atoi(defaultThemeIDStr)
		if err == nil {
			defaultTheme, err = themeServices.GetThemeByID(defaultThemeID)
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
			if defaultTheme.Type != "all" {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		}
	}

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
		"defaultTheme":  defaultTheme,
	}
	utils.JsonSuccessResponse(c, response)

}
