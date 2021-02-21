package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wejh-go/conf"
	"wejh-go/database"
	"wejh-go/routers"
)

func main() {
	conf.Init()     // 读取配置配置
	database.Init() // 初始化数据库

	// 将路由绑定到服务器引擎上
	r := gin.Default()
	routers.MiscRoutersInit(r)
	routers.UserRoutersInit(r)

	// 启动服务器
	err := r.Run()
	if err != nil {
		fmt.Printf("启动失败了，错误信息: %v \n", err)
	}
}
