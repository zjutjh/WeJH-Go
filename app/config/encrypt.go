package config

const encryptKey = "encryptKey"

func SetEncryptKey(value string) error {
	return setConfig(encryptKey, value)
}

func GetEncryptKey() string {
	return getConfig(encryptKey)
}

func IsSetEncryptKey() bool {
	return checkConfig(encryptKey)
}

func DelEncryptKey() error {
	return delConfig(encryptKey)
}
