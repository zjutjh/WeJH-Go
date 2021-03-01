package exception

const (
	ConfigNotFind         = "Can't find config.yaml, please add it before start.\n"
	ConfigError           = "Find config.yaml, but it's wrong. please check it.\n"
	DatabaseConnectFailed = "Try Connect DB Failed.\n"
	ServerStartFailed     = "Try Start Http Server Failed.\n"
)
