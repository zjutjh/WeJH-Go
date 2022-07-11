package config

import (
	"time"
)

const termYearKey = "termYearKey"
const termKey = "termKey"
const termStartDate = "termStartDate"

func SetTermInfo(yearValue, termValue string, termStartDateValue time.Time) error {
	err := setConfig(termYearKey, yearValue)
	if err != nil {
		return err
	}
	err = setConfig(termKey, termValue)
	if err != nil {
		return err
	}
	err = setConfig(termStartDate, termStartDateValue.String())
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
