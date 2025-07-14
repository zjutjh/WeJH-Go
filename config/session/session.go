package session

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(r *gin.Engine) {
	config := getConfig()
	switch config.Driver {
	case string(Redis):
		setRedis(r, config.Name)
		break
	case string(Memory):
		setMemory(r, config.Name)
		break
	default:
		zap.L().Fatal("Config Error")
	}
}
