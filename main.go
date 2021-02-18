package main

import (
	"github.com/gin-gonic/gin"
	"wejh-go/conf"
)

func main() {
	conf.Init() // 初始化配置
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
