package main

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"wejh-go/app/midwares"
	"wejh-go/app/utils/circuitBreaker"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		circuitBreaker.Probe.Start(ctx)
	}()

	<-ctx.Done()
	log.Println("shutting down gracefully")
	wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
