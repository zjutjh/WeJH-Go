package redis

import "wejh-go/config"

type Config struct {
	Host     string
	Port     string
	DB       int
	Password string
}

var Info Config

func init() {
	Info.Host = "localhost"
	if config.Config.IsSet("redis.host") {
		Info.Host = config.Config.GetString("redis.host")
	}

	Info.Port = "6379"
	if config.Config.IsSet("redis.port") {
		Info.Port = config.Config.GetString("redis.port")
	}
	Info.DB = 0
	if config.Config.IsSet("redis.db") {
		Info.DB = config.Config.GetInt("redis.db")
	}

	if config.Config.IsSet("redis.password") {
		Info.DB = config.Config.GetInt("redis.password")
	}

}
