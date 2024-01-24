package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func genTermForm(u *models.User, year, term, loginType string) url.Values {
	var password string

	if loginType == "OAUTH" {
		password = u.OauthPassword
	} else {
		password = u.ZFPassword
	}

	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", password)
	form.Add("type", loginType)
	form.Add("year", year)
	form.Add("term", term)
	return form
}

func GetClassTable(u *models.User, year, term, loginType string) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, funnelApi.ZFClassTable)
}

func GetScore(u *models.User, year, term, loginType string) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, funnelApi.ZFScore)
}

func GetMidTermScore(u *models.User, year, term, loginType string) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, funnelApi.ZFMidTermScore)
}

func GetExam(u *models.User, year, term, loginType string) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, funnelApi.ZFExam)
}

func GetRoom(u *models.User, year, term, campus, weekday, week, sections, loginType string) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	form.Add("campus", campus)
	form.Add("weekday", weekday)
	form.Add("week", week)
	form.Add("sections", sections)
	return FetchHandleOfPost(form, funnelApi.ZFRoom)
}

func BindPassword(u *models.User, year, term, loginType string) (interface{}, error) {
	var password string
	if loginType == "ZF" {
		password = u.ZFPassword
	} else if loginType == "OAUTH" {
		password = u.OauthPassword
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", password)
	form.Add("type", loginType)
	form.Add("year", year)
	form.Add("term", term)
	return FetchHandleOfPost(form, funnelApi.ZFExam)
}
