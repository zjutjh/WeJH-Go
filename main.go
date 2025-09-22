package main

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"wejh-go/app/midwares"
	"wejh-go/app/utils/circuitBreaker"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	srv := &http.Server{
		Addr:              ":" + config.Config.GetString("server.port"),
		Handler:           r,
		ReadHeaderTimeout: 2 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("Server Error Occurred", zap.Error(err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		circuitBreaker.Probe.Start(ctx)
	}()

	<-ctx.Done()
	zap.L().Info("Shutdown Server...")
	wg.Wait()

	// 关闭服务器（5秒超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown Failed", zap.Error(err))
	}

	zap.L().Info("Server Closed")
}
