package misc_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wejh-go/conf"
)

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
