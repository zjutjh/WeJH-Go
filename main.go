package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
	
	"wejh-go/app/utils/circuitBreaker"
	"wejh-go/config/database"
	"wejh-go/config/logger"
	"wejh-go/config/wechat"
	"wejh-go/register"
	"wejh-go/register/router"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"github.com/zjutjh/mygo/foundation/command"
	"github.com/zjutjh/mygo/foundation/httpserver"
)

func main() {

	command.Execute(
		register.Boot,    // 应用引导注册器
		register.Command, // 应用命令注册器
		func(cmd *cobra.Command, args []string) error {
			if err := logger.Init(); err != nil {
				log.Fatal(err.Error())
			}
			database.Init()
			circuitBreaker.Init()
	   		wechat.Init()
			ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()
			wg := &sync.WaitGroup{}

				// 启动HTTP Server
			wg.Add(1)
			go func() {
				defer wg.Done()
				httpserver.StartHTTPServer(router.Route)
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				circuitBreaker.Probe.Start(ctx)
			}()

				// 如有需要 可以额外启动其他服务
			<-ctx.Done()
			zap.L().Info("Shutdown Server...")
			wg.Wait()
			return nil
		},
	)


}
