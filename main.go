package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/app/midwares"
	"wejh-go/config/database"
	"wejh-go/config/router"
	"wejh-go/config/session"
)

func main() {
	database.Init()
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	session.Init(r)
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
