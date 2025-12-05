package cron

import (
	"github.com/zjutjh/mygo/config"
	"github.com/zjutjh/mygo/nlog"
)

type CronJob struct{}

func (CronJob) Run() {
	// TODO: 在此处编写定时任务业务逻辑
	nlog.Pick().WithField("app", config.AppName()).Debug("定时任务运行")
}
