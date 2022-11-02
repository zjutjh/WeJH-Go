package config

const schoolBusUrlKey = "schoolBusUrlKey"

func GetSchoolBusUrl() string {
	return getConfig(schoolBusUrlKey)
}

func SetSchoolBusUrl(url string) error {
	return setConfig(schoolBusUrlKey, url)
}
