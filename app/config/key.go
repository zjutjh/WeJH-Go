package config

const schoolBusUrlKey = "schoolBusUrlKey"

const webpUrlKey = "jpgUrlKey"

func GetSchoolBusUrl() string {
	return getConfig(schoolBusUrlKey)
}

func SetSchoolBusUrl(url string) error {
	return setConfig(schoolBusUrlKey, url)
}

func GetWebpUrlKey() string {
	return getConfig(webpUrlKey)
}

func SetWebpUrlKey(url string) error {
	return setConfig(webpUrlKey, url)
}

func GetNoticeBackGround(name string) string {
	return getConfig(name)
}
