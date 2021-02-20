package controllers

import (
	"fmt"
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
		"errcode":  1,
		"errmsg":   "获取时间成功",
		"redirect": nil,
		"data": gin.H{
			"day":      time.Now().Weekday(),
			"is_begin": week > 0,
			"month":    fmt.Sprintf("%d", time.Now().Month()),
			"term":     currentTerm,
			"week":     week,
		},
	})
}

func AnnouncementController(c *gin.Context) {
	var announcement map[string]interface{}
	err := conf.Config.UnmarshalKey("announcement", &announcement)
	if err != nil {
		panic(fmt.Errorf("配置读取失败, 请检查配置格式！\n %v", err))
	}
	c.JSON(http.StatusOK, gin.H{
		"data":     announcement,
		"errcode":  1,
		"errmsg":   "ok",
		"redirect": nil,
	})
}

func AppListController(c *gin.Context) {
	var appList []map[string]interface{}
	var icons map[string]interface{}
	err := conf.Config.UnmarshalKey("applicationsList", &appList)
	if err != nil {
		panic(fmt.Errorf("配置读取失败, 请检查配置格式！\n %v", err))
	}
	err = conf.Config.UnmarshalKey("icons", &icons)
	if err != nil {
		panic(fmt.Errorf("配置读取失败, 请检查配置格式！\n %v", err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"app-list": appList,
			"icons":    icons,
		},
		"errcode":  1,
		"errmsg":   "ok",
		"redirect": nil,
	})
}
