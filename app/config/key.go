package config

const schoolBusUrlKey = "schoolBusUrlKey"

const webpUrlKey = "jpgUrlKey"

const fileUrlKey = "fileUrlKey"

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

func GetFileUrlKey() string {
	return getConfig(fileUrlKey)
}

func SetFileUrlKey(url string) error {
	return setConfig(fileUrlKey, url)
}
