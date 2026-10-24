package config

const termYearKey = "termYearKey"
const termKey = "termKey"
const termStartDate = "termStartDate"
const scoreTermYearKey = "scoreTermYearKey"
const scoreTermKey = "scoreTermKey"

func SetSystemInfo(yearValue, termValue, termStartDateValue, scoreYearValue, scoreTermValue, jpgUrlValue, fileUrlValue, registerTipsValue, schoolBusUrlValue string) error {
	err := setConfig(termYearKey, yearValue)
	if err != nil {
		return err
	}
	err = setConfig(termKey, termValue)
	if err != nil {
		return err
	}
	err = setConfig(termStartDate, termStartDateValue)
	if err != nil {
		return err
	}
	err = setConfig(scoreTermYearKey, scoreYearValue)
	if err != nil {
		return err
	}
	err = setConfig(scoreTermKey, scoreTermValue)
	if err != nil {
		return err
	}
	err = setConfig(schoolBusUrlKey, schoolBusUrlValue)
	if err != nil {
		return err
	}
	err = setConfig(webpUrlKey, jpgUrlValue)
	if err != nil {
		return err
	}
	err = setConfig(fileUrlKey, fileUrlValue)
	if err != nil {
		return err
	}
	err = setConfig(registerTipsKey, registerTipsValue)
	if err != nil {
		return err
	}
	return err
}


func GetTermInfo() (year, term, startDate, scoreYear, scoreTerm string) {
	return getConfig(termYearKey), getConfig(termKey), getConfig(termStartDate),
		getConfig(scoreTermYearKey), getConfig(scoreTermKey)
}

func GetScoreTermInfo() (scoreYear, scoreTerm string) {
	return getConfig(scoreTermYearKey), getConfig(scoreTermKey)
}

func IsSetTermInfo() bool {
	return checkConfig(termYearKey) && checkConfig(termKey) && checkConfig(termStartDate) &&
		checkConfig(scoreTermYearKey) && checkConfig(scoreTermKey)
}

func DelTermInfo() []error {
	var result []error
	result = append(result,
		delConfig(termYearKey),
		delConfig(termKey),
		delConfig(termStartDate),
		delConfig(scoreTermYearKey),
		delConfig(scoreTermKey),
	)
	return result
}
