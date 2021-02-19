package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wejh-go/conf"
	"wejh-go/routers"
)

func main() {
	conf.Init() // 初始化配置
	r := gin.Default()
	routers.MiscRoutersInit(r) // 将路由绑定到服务器引擎上
	err := r.Run(":8090")
	if err != nil {
		fmt.Printf("启动失败了，错误信息: %v \n", err)
	}
}
