package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func setMemory(r *gin.Engine, name string) {
	store := memstore.NewStore()
	r.Use(sessions.Sessions(name, store))
}
