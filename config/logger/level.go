package logger

import "go.uber.org/zap"

// Level 日志级别
type Level uint8

// 日志级别常量
const (
	LevelFatal  Level = 0
	LevelPanic  Level = 1
	LevelDpanic Level = 2
	LevelError  Level = 3
	LevelWarn   Level = 4
	LevelInfo   Level = 5
	LevelDebug  Level = 6
)

// GetLogFunc 根据日志级别返回对应的日志函数
func GetLogFunc(level Level) func(string, ...zap.Field) {
	// 创建日志级别映射表
	logMap := map[Level]func(string, ...zap.Field){
		LevelFatal:  zap.L().Fatal,
		LevelPanic:  zap.L().Panic,
		LevelDpanic: zap.L().DPanic,
		LevelError:  zap.L().Error,
		LevelWarn:   zap.L().Warn,
		LevelInfo:   zap.L().Info,
		LevelDebug:  zap.L().Debug,
	}

	// 根据日志级别记录日志
	if logFunc, ok := logMap[level]; ok {
		return logFunc
	}
	return zap.L().Info
}
