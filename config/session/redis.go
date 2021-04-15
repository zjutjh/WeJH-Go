package session

import (
	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"wejh-go/config/redis"
)

func setRedis(r *gin.Engine, name string) {
	store, _ := sessionRedis.NewStore(10, "tcp", redis.Info.Host+":"+redis.Info.Port, "", []byte("secret"))
	r.Use(sessions.Sessions(name, store))
}
