package cache

import (
	"log"
	"strings"
	"wejh-go/config"
	"wejh-go/exception"
)

type driver string

const (
	Memory driver = "memory"
	Redis  driver = "redis"
)

func init() {
	myDriver := string(Memory)
	if !config.Config.IsSet("cache.driver") {
		myDriver = config.Config.GetString("cache.driver")
	}
	myDriver = strings.ToLower(myDriver)
	switch myDriver {
	case string(Redis):
		break
	case string(Memory):
		break
	default:
		log.Fatal(exception.ConfigError)
	}

}
