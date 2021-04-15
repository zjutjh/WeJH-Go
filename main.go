package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/config/database"
	"wejh-go/config/router"
	"wejh-go/config/session"
	"wejh-go/exception"
)

func main() {
	database.Init()

	r := gin.Default()

	session.Init(r)
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal(exception.ServerStartFailed, err)
	}
}
