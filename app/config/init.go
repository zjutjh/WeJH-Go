package config

const initKey = "initKey"

func SetInit() error {
	return setConfig(initKey, "True")
}

func ResetInit() error {
	return setConfig(initKey, "False")
}

func GetInit() bool {
	return getConfig(initKey) == "True"
}
