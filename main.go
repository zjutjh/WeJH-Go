package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/config"
	"wejh-go/exception"
	"wejh-go/service/database"
	"wejh-go/service/router"
)

func main() {
	config.Init()
	database.Init()

	r := gin.Default()
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal(exception.ServerStartFailed, err)
	}
}
