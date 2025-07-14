package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/app/midwares"
	"wejh-go/app/utils/server"
	"wejh-go/config/config"
	"wejh-go/config/database"
	"wejh-go/config/logger"
	"wejh-go/config/router"
	"wejh-go/config/session"
)

func main() {
	if err := logger.Init(); err != nil {
		log.Fatal(err.Error())
	}
	database.Init()
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	session.Init(r)
	router.Init(r)

	server.Run(r, ":"+config.Config.GetString("server.port"))
}
