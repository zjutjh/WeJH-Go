package logger

import (
	"github.com/zjutjh/WeJH-SDK/zapHelper"
	"github.com/zjutjh/mygo/config"
	"go.uber.org/zap"
)

func Init() error {
	zapInfo := zapHelper.InfoConfig{
		StacktraceLevel:   "warn",
		DisableStacktrace: config.Pick().GetBool("log.disableStacktrace"), // 是否禁用堆栈跟踪
		ConsoleLevel:      "Info",           // 日志级别
		Name:              config.Pick().GetString("log.name"),            // 日志名称
		Writer:            config.Pick().GetString("log.writer"),          // 日志输出方式
		LoggerDir:         config.Pick().GetString("log.loggerDir"),       // 日志目录
		LogCompress:       config.Pick().GetBool("log.logCompress"),       // 是否压缩日志
		LogMaxSize:        config.Pick().GetInt("log.logMaxSize"),         // 日志文件最大大小（单位：MB）
		LogMaxAge:         config.Pick().GetInt("log.logMaxAge"),          // 日志保存天数
	}
	logger, err := zapHelper.Init(&zapInfo)
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	zap.L().Info("Logger initialized")
	return nil
}
