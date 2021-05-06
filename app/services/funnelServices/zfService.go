package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
	"wejh-go/errors"
)

func genTermForm(u *models.User, year, term string) url.Values {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.ZFPassword)
	form.Add("year", year)
	form.Add("term", term)
	return form
}

func GetClassTable(u *models.User, year, term string) (interface{}, error) {
	if u.ZFPassword == "" {
		return nil, errors.PasswordWrong
	}
	form := genTermForm(u, year, term)
	return FetchHandleOfPost(form, funnelApi.ZFClassTable)
}

func GetScore(u *models.User, year, term string) (interface{}, error) {
	if u.ZFPassword == "" {
		return nil, errors.PasswordWrong
	}
	form := genTermForm(u, year, term)
	return FetchHandleOfPost(form, funnelApi.ZFScore)
}

func GetExam(u *models.User, year, term string) (interface{}, error) {
	if u.ZFPassword == "" {
		return nil, errors.PasswordWrong
	}
	form := genTermForm(u, year, term)
	return FetchHandleOfPost(form, funnelApi.ZFExam)
}

func GetRoom(u *models.User, year, term, campus, weekday, week, sections string) (interface{}, error) {
	if u.ZFPassword == "" {
		return nil, errors.PasswordWrong
	}
	form := genTermForm(u, year, term)
	form.Add("campus", campus)
	form.Add("weekday", weekday)
	form.Add("week", week)
	form.Add("sections", sections)
	return FetchHandleOfPost(form, funnelApi.ZFRoom)
}
