package logger

import (
	"github.com/zjutjh/WeJH-SDK/zapHelper"
	"go.uber.org/zap"
	"wejh-go/config/config"
)

func Init() error {
	zapInfo := zapHelper.InfoConfig{
		StacktraceLevel:   "warn",
		DisableStacktrace: config.Config.GetBool("log.disableStacktrace"), // 是否禁用堆栈跟踪
		ConsoleLevel:      config.Config.GetString("log.level"),           // 日志级别
		Name:              config.Config.GetString("log.name"),            // 日志名称
		Writer:            config.Config.GetString("log.writer"),          // 日志输出方式
		LoggerDir:         config.Config.GetString("log.loggerDir"),       // 日志目录
		LogCompress:       config.Config.GetBool("log.logCompress"),       // 是否压缩日志
		LogMaxSize:        config.Config.GetInt("log.logMaxSize"),         // 日志文件最大大小（单位：MB）
		LogMaxAge:         config.Config.GetInt("log.logMaxAge"),          // 日志保存天数
	}
	logger, err := zapHelper.Init(&zapInfo)
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	zap.L().Info("Logger initialized")
	return nil
}
