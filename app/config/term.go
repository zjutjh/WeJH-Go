package config

const termYearKey = "termYearKey"
const termKey = "termKey"
const termStartDate = "termStartDate"

func SetTermInfo(yearValue, termValue, termStartDateValue string) error {
	err := setConfig(termYearKey, yearValue)
	if err != nil {
		return err
	}
	err = setConfig(termKey, termValue)
	if err != nil {
		return err
	}
	err = setConfig(termStartDate, termStartDateValue)
	return err
}

func GetTermInfo() (string, string, string) {
	return getConfig(termYearKey), getConfig(termKey), getConfig(termStartDate)
}

func IsSetTermInfo() bool {
	return checkConfig(termYearKey) && checkConfig(termKey) && checkConfig(termStartDate)
}

func DelTermInfo() []error {
	var result []error
	errTermYear := delConfig(termYearKey)
	errTerm := delConfig(termKey)
	errStartDate := delConfig(termStartDate)
	return append(result, errTermYear, errTerm, errStartDate)
}
