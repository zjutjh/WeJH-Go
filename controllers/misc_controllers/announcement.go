package misc_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wejh-go/conf"
)

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
