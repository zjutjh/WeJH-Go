package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wejh-go/conf"
)

func TimeController(c *gin.Context) {
	startTimeString := conf.Config.GetString("term.term_start_date")
	currentTerm := conf.Config.GetString("term.current_term") // 当前学期
	startTime, _ := time.Parse("2006-01-02", startTimeString) // 学期开始的时间
	currentTime := time.Now()

	// 计算第几周
	week := (currentTime.Unix()-startTime.Unix())/3600/24/7 + 1

	c.JSON(http.StatusOK, gin.H{
		"errcode": 1,
		"errmsg":  "获取时间成功",
		"data": gin.H{
			"day":      time.Now().Weekday(),
			"is_begin": week > 0,
			"month":    time.Now().Month(),
			"term":     currentTerm,
			"week":     week,
		},
	})
}
