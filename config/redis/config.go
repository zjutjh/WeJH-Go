package redis

import "wejh-go/config/config"

func getConfig() redisConfig {
	Info := redisConfig{
		Host:     "localhost",
		Port:     "6379",
		DB:       0,
		Password: "",
	}
	if config.Config.IsSet("redis.host") {
		Info.Host = config.Config.GetString("redis.host")
	}
	if config.Config.IsSet("redis.port") {
		Info.Port = config.Config.GetString("redis.port")
	}
	if config.Config.IsSet("redis.db") {
		Info.DB = config.Config.GetInt("redis.db")
	}
	if config.Config.IsSet("redis.pass") {
		Info.Password = config.Config.GetString("redis.pass")
	}
	return Info
}
