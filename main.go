package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"wejh-go/app/utils/circuitBreaker"
	"wejh-go/register"
	"wejh-go/register/router"

	"github.com/spf13/cobra"
	"github.com/zjutjh/mygo/foundation/command"
	"github.com/zjutjh/mygo/foundation/crontab"
	"github.com/zjutjh/mygo/foundation/httpserver"
)

func main() {

	command.Execute(
		register.Boot,    // 应用引导注册器
		register.Command, // 应用命令注册器
		func(cmd *cobra.Command, args []string) error {
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
				circuitBreaker.Probe.Start(ctx)//circuitBreaker 后面放到 cron
			}()

			// 启动HTTP Server伴生定时任务
			wg.Add(1)
			go func() {
				defer wg.Done()
				crontab.Run(register.CronWithHTTPServer)
			}()

			wg.Wait()
			return nil
		},
	)

}
