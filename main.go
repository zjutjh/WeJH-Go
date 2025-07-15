package main

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wejh-go/app/midwares"
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

	srv := &http.Server{
		Addr:              ":" + config.Config.GetString("server.port"),
		Handler:           r,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// 启动服务器协程
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("Server Error Occurred", zap.Error(err))
		}
	}()

	// 阻塞并监听结束信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutdown Server...")

	// 关闭服务器（5秒超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown Failed", zap.Error(err))
	}

	zap.L().Info("Server Closed")
}
