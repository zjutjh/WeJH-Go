package config

const encryptKey = "encryptKey"

func SetEncryptKey(value string) error {
	return setConfig(encryptKey, value)
}
func GetEncryptKey() string {
	return getConfig(encryptKey)
}
