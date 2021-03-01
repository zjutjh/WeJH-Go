package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/exception"
	"wejh-go/service/database"
	"wejh-go/service/router"
	"wejh-go/service/session"
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
