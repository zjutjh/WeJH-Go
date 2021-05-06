package session

import (
	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type redisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func setRedis(r *gin.Engine, name string) {
	Info := getRedisConfig()
	store, _ := sessionRedis.NewStore(10, "tcp", Info.Host+":"+Info.Port, "", []byte("secret"))
	r.Use(sessions.Sessions(name, store))
}
