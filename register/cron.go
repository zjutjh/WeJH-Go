package register

import (
	cron2 "github.com/robfig/cron/v3"
)

func CronWithHTTPServer(c *cron2.Cron) {
	// 定时任务 (默认随HTTP Server伴生运行)
	// c.AddJob("* * * * * *", cron.XXXJob{})
}

func Cron(c *cron2.Cron) {
	// 业务定时任务 (独立运行)
	// c.AddJob("* * * * * *", cron.XXXJob{})
}
